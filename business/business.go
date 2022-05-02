package business

import (
	"log"
	"sync"

	"github.com/pavelerokhin/go-and-scrape/business/modules"
	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"github.com/pavelerokhin/go-and-scrape/storage"
)

type Business struct {
	scrapper *modules.Scrapper
	storage  storage.Storage
	logger   *log.Logger
}

func GetBusinessLogic(logger *log.Logger, storage storage.Storage) Business {
	return Business{scrapper: modules.NewScrapper(logger), storage: storage, logger: logger}
}

func (b *Business) ScrapeAndPersist(mediumConfig configs.MediumConfig,
	wg *sync.WaitGroup) {
	defer wg.Done()

	// scrape
	medium := b.storage.GetMediumByURL(mediumConfig.URL)
	articlePreviewsWithArticle := b.scrapper.Scrap(&mediumConfig)

	// persist
	if len(articlePreviewsWithArticle) > 0 {
		articlePreviewsWithArticle = modules.NormalizeText(articlePreviewsWithArticle)

		if medium == nil || medium.URL == "" {
			b.storage.SaveMedium(&entities.Medium{
				Name:            mediumConfig.Name,
				URL:             mediumConfig.URL,
				ArticlePreviews: articlePreviewsWithArticle,
			})
		} else {
			for _, articlePreview := range articlePreviewsWithArticle {
				a := articlePreview
				a.MediumID = medium.ID
				b.storage.SaveArticle(&a)
			}
		}
	}
}
