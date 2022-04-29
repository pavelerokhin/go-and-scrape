package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/pavelerokhin/go-and-scrape/business"
	"github.com/pavelerokhin/go-and-scrape/storage"
)

var (
	logger = log.New(os.Stdout, "go-and-scrape-logger", log.LstdFlags|log.Llongfile)
	wg     sync.WaitGroup
)

func main() {
	fileConfig, err := business.ReadMediumConfig("medium-config.yaml")
	check(err)

	if len(fileConfig.Mediums) == 0 {
		fmt.Println("no mediums set")
		os.Exit(0)
	}
	articleStorage, err := storage.NewSQLiteArticleRepo(fileConfig.Mediums[0].MediumConfig.FileName)
	check(err)

	for _, medium := range fileConfig.Mediums {
		wg.Add(1)
		go business.ScrapeAndPersistWorker(articleStorage, logger, medium.MediumConfig, &wg)
	}
	wg.Wait()
}

func check(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
