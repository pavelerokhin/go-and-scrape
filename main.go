package main

import (
	"fmt"
	"github.com/pavelerokhin/go-and-scrape/config"
	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"log"
	"os"
	"sync"

	"github.com/pavelerokhin/go-and-scrape/business"
	"github.com/pavelerokhin/go-and-scrape/storage"
)

var (
	businessLogic business.Business
	repo          storage.Storage
	wg            sync.WaitGroup
)

func main() {
	logger := log.New(os.Stdout, "go-and-scrape ", log.LstdFlags|log.Lshortfile)
	configFile, err := config.ReadMediumConfig("medium-config.yaml")
	check(err)
	checkConfig(configFile)
	repo, err = storage.NewSQLiteArticleRepo(configFile.Mediums[0].MediumConfig.FileName,
		logger)
	check(err)
	businessLogic = business.GetBusinessLogic(logger, repo)

	for _, medium := range configFile.Mediums {
		wg.Add(1)
		go businessLogic.ScrapeAndPersist(medium.MediumConfig, &wg)
	}
	wg.Wait()
}

func checkConfig(configFile *configs.ConfigFile) {
	if len(configFile.Mediums) == 0 {
		fmt.Println("no mediums set")
		os.Exit(0)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
