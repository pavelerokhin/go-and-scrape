package main

import (
	"fmt"
	"github.com/pavelerokhin/go-and-scrape/scrapper/business"
	"github.com/pavelerokhin/go-and-scrape/scrapper/config"
	"github.com/pavelerokhin/go-and-scrape/scrapper/storage"
	"log"
	"os"
)

var (
	businessLogic business.Business
	repo          storage.Storage
)

func main() {
	logger := log.New(os.Stdout, "go-and-scrape ", log.LstdFlags|log.Lshortfile)

	mediaConfig, persistenceConfig, err := config.ReadConfig("../config/config.yaml")
	check(err)
	config.CheckMediaConfig(&mediaConfig)
	repo, err = storage.NewSQLiteArticleRepo(persistenceConfig.Filename, logger)
	check(err)
	businessLogic = business.GetBusinessLogic(logger, repo)

	// scheduler
	if persistenceConfig.Interval != 0 {
		businessLogic.ScheduleWithInterval(mediaConfig, persistenceConfig)
	} else {
		businessLogic.ScheduleOnce(mediaConfig)
	}
}

func check(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
