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
	mediaConfig, persistenceConfig, err := config.ReadConfig("config.yaml")
	check(err)
	checkMedia(&mediaConfig)
	repo, err = storage.NewSQLiteArticleRepo(persistenceConfig.Filename, logger)
	check(err)
	businessLogic = business.GetBusinessLogic(logger, repo)

	for _, medium := range mediaConfig {
		wg.Add(1)
		go businessLogic.ScrapeAndPersist(medium.MediumConfig, &wg)
	}
	wg.Wait()
}

func checkMedia(mediaConfig *configs.MediumConfigs) {
	if len(*mediaConfig) == 0 {
		fmt.Println("no media set")
		os.Exit(0)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
