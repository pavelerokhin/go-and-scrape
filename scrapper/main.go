package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pavelerokhin/go-and-scrape/business"
	"github.com/pavelerokhin/go-and-scrape/config"
	"github.com/pavelerokhin/go-and-scrape/storage"
)

var (
	businessLogic business.Business
	repo          storage.Storage
)

func main() {
	logger := log.New(os.Stdout, "go-and-scrape ", log.LstdFlags|log.Lshortfile)
	mediaConfig, persistenceConfig, err := config.ReadConfig("config/config.yaml")
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
		fmt.Println(err)
		os.Exit(1)
	}
}
