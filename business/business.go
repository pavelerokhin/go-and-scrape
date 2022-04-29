package business

import (
	"github.com/pavelerokhin/go-and-scrape/business/modules"
	"log"
	"sync"

	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"github.com/pavelerokhin/go-and-scrape/storage"
)

type Business struct {
	logger *log.Logger
}

func GetBusinessLogic(logger *log.Logger) *Business {
	return &Business{logger: logger}
}

func (b *Business) ScrapeAndPersist(storage *storage.SQLiteRepo, mediumConfig *configs.MediumConfig, wg *sync.WaitGroup) error {
	defer wg.Done()

	articles, err := modules.ScrapMedium(mediumConfig)
	if err != nil {
		return err
	}

	var medium *entities.Medium
	medium, err = storage.GetMediumByURL(mediumConfig.URL)
	if err != nil {
		return err
	}

	if len(articles) > 0 {
		articles = modules.NormalizeArticlesNLP(articles)

		if medium.URL == "" {
			_, err = storage.SaveMedium(&entities.Medium{
				Name:     mediumConfig.Name,
				URL:      mediumConfig.URL,
				Articles: articles,
			})
		} else {
			for _, article := range articles {
				a := article
				a.MediumID = medium.ID
				_, err = storage.SaveArticle(&a)
			}
		}
	}
	if err != nil {
		return err
	}

	return nil
}
