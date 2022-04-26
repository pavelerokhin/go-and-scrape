package business

import (
	"github.com/pavelerokhin/go-and-scrape/models"
	"sync"
)

func ScrapeAndPersist(medium *models.Medium, wg *sync.WaitGroup) error {
	defer wg.Done()

	articles, err := ScrapMedium(medium)
	if err != nil {
		return err
	}

	if len(articles) > 0 {
		err = WriteCSV(articles, medium)
	}
	if err != nil {
		return err
	}

	return nil
}
