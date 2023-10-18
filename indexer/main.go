package main

import (
	"flag"
	"indexador/db"
	"indexador/filesearch"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "cpuProf", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "memProf", "write memory profile to `file`")

func main() {
	// dataDirectory, err := GetDataDirectory()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("No se pudo crear el CPU Profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("No se pudo crear el CPU Profile: ", err)
		}
		defer pprof.StopCPUProfile()

		if *memprofile != "" {
			f, err := os.Create(*memprofile)
			if err != nil {
				log.Fatal("No se pudo crear el Memory Profile: ", err)
			}
			defer f.Close()
			runtime.GC()
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("No se pudo crear el Memory Profile: ", err)
			}
		}

		mailExplorer := new(filesearch.MailIndexer)
		mailExplorer.DbProcesor = new(db.ZincSearch)
		mailExplorer.RootDirectory = "L:\\Proyectos\\Proyectos_go\\enron_mail_20110402"

		err = mailExplorer.IndexFilesToDB()

		if err != nil {
			log.Println("Error")
		}
	}
}

func GetDataDirectory() (string, error) {
	if len(os.Args) > 2 {
		return os.Args[1], nil
	} else {
		ex, err := os.Executable()
		if err != nil {
			return "", err
		}
		return filepath.Dir(ex), nil
	}
}
