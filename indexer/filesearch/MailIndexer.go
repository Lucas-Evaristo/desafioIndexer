package filesearch

import (
	"indexador/db"
	"indexador/model"
	"io/fs"
	"log"
	"path/filepath"
	"sync"
)

var chunkSize = 15000

type MailIndexer struct {
	mailParser    MailParser
	DbProcesor    db.Database
	RootDirectory string
}

func (m MailIndexer) IndexFilesToDB() error {
	BulkEmail := new(model.BulkEmail)
	BulkEmail.Index = "Emails2"
	parsed := make(chan map[string]string)
	wg := new(sync.WaitGroup)

	files, err := m.obtainFilesInChunks(chunkSize)
	if err != nil {
		return err
	}

	for _, fs := range files {
		wg.Add(1)
		go func(files []string) {
			defer wg.Done()
			for _, file := range files {
				convertedEmail, err := m.mailParser.ConvertFileToEmail(file)
				if err != nil {
					log.Println("Error al procesar archivo: " + file)
					continue
				}
				parsed <- convertedEmail
			}
		}(fs)
	}

	go func() {
		defer close(parsed)
		wg.Wait()
	}()

	for p := range parsed {
		BulkEmail.Emails = append(BulkEmail.Emails, p)
	}

	err = m.DbProcesor.InsertRecords(*BulkEmail)
	if err != nil {
		return err
	}

	return nil
}

func (m *MailIndexer) obtainFilesInChunks(maxSize int) ([][]string, error) {
	var files [][]string
	var chunk []string

	err := filepath.WalkDir(m.RootDirectory, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			log.Println("Error al procesar directorio: " + info.Name())
			return err
		}
		if !info.IsDir() {
			chunk = append(chunk, path)
			if len(chunk) >= maxSize {
				files = append(files, chunk)
				chunk = nil
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(chunk) > 0 {
		files = append(files, chunk)
	}

	return files, nil
}
