package main

import (
	"fmt"
	"github.com/pavelerokhin/go-and-scrape/business/modules"
	"os"
	"sync"

	"github.com/pavelerokhin/go-and-scrape/business"
	"github.com/pavelerokhin/go-and-scrape/storage"
)

var (
	wg sync.WaitGroup
)

func main() {
	fileConfig, err := modules.ReadMediumConfig("medium-config.yaml")
	check(err)

	if len(fileConfig.Mediums) == 0 {
		fmt.Println("no mediums set")
		os.Exit(0)
	}
	articleStorage, err := storage.NewSQLiteArticleRepo(fileConfig.Mediums[0].MediumConfig.FileName)
	check(err)

	for _, medium := range fileConfig.Mediums {
		wg.Add(1)
		m := medium.MediumConfig
		go func() {
			err := business.ScrapeAndPersist(articleStorage, &m, &wg)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
