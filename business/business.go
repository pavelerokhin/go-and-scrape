package business

import (
	"io/ioutil"
	"log"
	"sync"

	"github.com/pavelerokhin/go-and-scrape/business/modules"
	"gopkg.in/yaml.v2"

	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"github.com/pavelerokhin/go-and-scrape/storage"
)

type Business struct {
	scrapper *modules.Scrapper
	logger   *log.Logger
}

func GetBusinessLogic(logger *log.Logger) *Business {
	return &Business{scrapper: modules.NewScrapper(logger), logger: logger}
}

// ReadMediumConfig reads `configs` file and returns list of Mediums
func (b *Business) ReadMediumConfig(configFilePath string) (*configs.ConfigFile, error) {
	b.logger.Printf("reading the config file %s\n", configFilePath)
	bytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		b.logger.Printf("error reading the config file %e\n", err)
		return nil, err
	}

	configFile := &configs.ConfigFile{}
	err = yaml.Unmarshal(bytes, configFile)
	if err != nil {
		b.logger.Printf("error reading the config file %e\n", err)
		return nil, err
	}

	return configFile, nil
}

func (b *Business) ScrapeAndPersist(storage *storage.SQLiteRepo, mediumConfig configs.MediumConfig,
	wg *sync.WaitGroup) {
	defer wg.Done()

	articles := b.scrapper.ScrapMedium(&mediumConfig)
	medium := storage.GetMediumByURL(mediumConfig.URL)

	if len(articles) > 0 {
		articles = modules.NormalizeArticlesNLP(articles)

		if medium == nil || medium.URL == "" {
			storage.SaveMedium(&entities.Medium{
				Name:     mediumConfig.Name,
				URL:      mediumConfig.URL,
				Articles: articles,
			})
		} else {
			for _, article := range articles {
				a := article
				a.MediumID = medium.ID
				storage.SaveArticle(&a)
			}
		}
	}
}
