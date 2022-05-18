package business

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/pavelerokhin/go-and-scrape/scrapper/business/modules"
	"github.com/pavelerokhin/go-and-scrape/scrapper/models/configs"
	"github.com/pavelerokhin/go-and-scrape/scrapper/models/entities"
	"github.com/pavelerokhin/go-and-scrape/scrapper/storage"
)

type Business struct {
	scrapper *modules.Scrapper
	storage  storage.Storage
	logger   *log.Logger
}

func GetBusinessLogic(logger *log.Logger, storage storage.Storage) Business {
	return Business{scrapper: modules.NewScrapper(logger), storage: storage, logger: logger}
}

func (b *Business) ScheduleOnce(mediaConfig configs.MediaConfig) {
	b.logger.Printf("scheduling scrape and persist of %d media at %v", len(mediaConfig),
		time.Now())
	var wg sync.WaitGroup
	for _, medium := range mediaConfig {
		wg.Add(1)
		go b.scrapeNLPPersist(medium.MediumConfig, &wg)
	}
	wg.Wait()
}

func (b *Business) ScheduleWithInterval(mediaConfig configs.MediaConfig,
	persistenceConfig *configs.PersistenceConfig) {
	b.logger.Printf("scheduling scrape and persist work every %d sec",
		persistenceConfig.Interval)
	for {
		go b.ScheduleOnce(mediaConfig)
		time.Sleep(persistenceConfig.Interval * time.Second)
	}
}

func (b *Business) scrapeNLPPersist(mediaConfig configs.MediumConfig, wg *sync.WaitGroup) {
	defer wg.Done()

	// scrape
	medium := b.storage.GetMediumByURL(mediaConfig.URL)
	articlePreviewsWithArticle := b.scrapper.Scrap(&mediaConfig)

	// persist
	if len(articlePreviewsWithArticle) > 0 {
		b.logger.Println("article's text normalization")
		articlePreviewsWithArticle = modules.Normalize(articlePreviewsWithArticle)

		if medium == nil || medium.URL == "" {
			b.storage.SaveMedium(&entities.Medium{
				Name:            mediaConfig.Name,
				URL:             mediaConfig.URL,
				ArticlePreviews: articlePreviewsWithArticle,
			})
		} else {
			for _, articlePreview := range articlePreviewsWithArticle {
				a := articlePreview
				a.MediumID = medium.ID
				if b.storage.GetArticleByURL(a.RelativeURL) == nil {
					b.storage.SaveArticle(&a)
				} else {
					b.logger.Printf("article with URL %s for medium %s is present in DB and will not be persisted",
						a.RelativeURL, medium.Name)
				}
			}
		}
	}
	// run Python nlp module
	b.logger.Println("run Named Entity Recognition module")
	nlpNerRequest()
}

func nlpNerRequest() {
	url := "http://127.0.0.1:7070/ner"

	_, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
}
