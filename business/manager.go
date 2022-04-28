package business

import (
	"sync"

	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"github.com/pavelerokhin/go-and-scrape/storage"
)

func ScrapeAndPersist(articleStorage *storage.SQLiteArticleRepo, mediumConfig *configs.MediumConfig, wg *sync.WaitGroup) error {
	defer wg.Done()

	articles, err := ScrapMedium(mediumConfig)
	if err != nil {
		return err
	}

	if len(articles) > 0 {
		medium := entities.Medium{
			Name:     mediumConfig.Name,
			URL:      mediumConfig.URL,
			Articles: articles,
		}

		_, err = articleStorage.Save(&medium)
	}
	if err != nil {
		return err
	}

	return nil
}
