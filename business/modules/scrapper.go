package modules

import (
	"fmt"
	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Scrapper struct {
	logger *log.Logger
}

func NewScrapper(logger *log.Logger) *Scrapper {
	return &Scrapper{logger: logger}
}

// ScrapMedium is a function that scpapa medium and returns a slice of Article
// objects and error
func (s *Scrapper) ScrapMedium(mediumConfig *configs.MediumConfig) []entities.Article {
	s.logger.Printf("start to scrap medium %s", mediumConfig.Name)
	response, err := http.Get(mediumConfig.URL)
	if err != nil {
		s.logger.Printf("error sending the request: %e\n", err)
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		s.logger.Printf("status code: %v\n", response.StatusCode)
		return nil
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		s.logger.Printf("error while reading from the document: %e\n", err)
		return nil
	}

	newsContainer := document.Find(mediumConfig.HTMLTags.Article)
	if newsContainer.Size() == 0 {
		s.logger.Println("no news found")
		return nil
	}

	s.logger.Printf("%d articles has been found for the medium %s\n", newsContainer.Size(),
		mediumConfig.Name)

	var articles []entities.Article
	newsContainer.Each(func(i int, item *goquery.Selection) {
		tag := strings.TrimSpace(item.Find(mediumConfig.HTMLTags.Tag).Text())
		title := strings.TrimSpace(item.Find(mediumConfig.HTMLTags.Title).Text())
		subtitle := strings.TrimSpace(item.Find(mediumConfig.HTMLTags.Subtitle).Text())
		urlArticle, _ := item.Find(mediumConfig.HTMLTags.URL).Attr("href")
		urlArticle = fmt.Sprintf("%s%s", mediumConfig.URL, urlArticle)

		articles = append(articles, entities.Article{
			Tag:      tag,
			Title:    title,
			Subtitle: subtitle,
			URL:      urlArticle,
		})
	})

	s.logger.Printf("scrapping of medium %s finished successfully\n", mediumConfig.Name)
	return articles
}
