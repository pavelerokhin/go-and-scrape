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
	wg sync.WaitGroup
)

func main() {
	logger := log.New(os.Stdout, "go-and-scrape-log", log.LstdFlags|log.Llongfile)
	businessLogic := business.GetBusinessLogic(logger)
	fileConfig, err := businessLogic.ReadMediumConfig("medium-config.yaml")
	check(err)

	if len(fileConfig.Mediums) == 0 {
		fmt.Println("no mediums set")
		os.Exit(0)
	}
	articleStorage, err := storage.NewSQLiteArticleRepo(fileConfig.Mediums[0].MediumConfig.FileName,
		logger)
	check(err)

	for _, medium := range fileConfig.Mediums {
		wg.Add(1)
		go businessLogic.ScrapeAndPersist(articleStorage, medium.MediumConfig, &wg)
	}
	wg.Wait()
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
