package filesearch

import (
	"bufio"
	"indexador/model"
	"io"
	"os"
	"strings"
)

type MailParser struct {
}

func (m MailParser) ConvertFileToEmail(path string) (map[string]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	emailFields := model.NewEmailFields()
	email := make(map[string]string)
	currentField := ""

	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			return nil, err
		}
		line = strings.TrimSpace(strings.TrimRight(line, "\r"))

		field, content, found := strings.Cut(line, ":")

		if found && currentField != "Body" {
			_, ok := emailFields[field]

			if ok {
				email[field] = content
				currentField = field
			} else {
				email[currentField] += content
			}
		} else {
			if line == "" {
				currentField = "Body"
			} else {
				email[currentField] += line
			}
		}

		if err == io.EOF {
			break
		}

	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return email, nil
}
