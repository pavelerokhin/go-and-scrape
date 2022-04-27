package business

import (
	"sync"

	"github.com/pavelerokhin/go-and-scrape/models"
	"github.com/pavelerokhin/go-and-scrape/store"
)

var (
	articleStorage *store.ArticleStorage
)

func ScrapeAndPersist(medium *models.Medium, wg *sync.WaitGroup) error {
	defer wg.Done()
	articleStorage, err := store.NewSQLiteArtcleRepo(medium)

	articles, err := ScrapMedium(medium)
	if err != nil {
		return err
	}

	if len(articles) > 0 {
		for _, article := range articles {
			_, err = articleStorage.Save(&article)
		}
	}
	if err != nil {
		return err
	}

	return nil
}
