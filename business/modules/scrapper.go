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

// ScrapMedium is a function that scraps medium and returns a slice of Article
// objects and error
func (s *Scrapper) ScrapMedium(mediumConfig *configs.MediumConfig) []entities.Article {
	s.logger.Printf("start to scrap medium %s", mediumConfig.Name)
	response, err := http.Get(mediumConfig.URL)
	if err != nil {
		s.logger.Printf("error sending the request: %e", err)
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		s.logger.Printf("status code: %v", response.StatusCode)
		return nil
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		s.logger.Printf("error while reading from the document: %e", err)
		return nil
	}

	newsContainer := document.Find(mediumConfig.HTMLTags.Article)
	if newsContainer.Size() == 0 {
		s.logger.Println("no news found")
		return nil
	}

	s.logger.Printf("%d articles has been found for the medium %s", newsContainer.Size(),
		mediumConfig.Name)

	var articles []entities.Article
	newsContainer.Each(func(i int, item *goquery.Selection) {
		tag := strings.TrimSpace(item.Find(mediumConfig.HTMLTags.Tag).Text())
		title := strings.TrimSpace(item.Find(mediumConfig.HTMLTags.Title).Text())
		subtitle := strings.TrimSpace(item.Find(mediumConfig.HTMLTags.Subtitle).Text())
		urlArticle, _ := item.Find(mediumConfig.HTMLTags.URL).Attr("href")

		urlArticle = s.getUrl(mediumConfig.URL, urlArticle)

		articles = append(articles, entities.Article{
			Tag:      tag,
			Title:    title,
			Subtitle: subtitle,
			URL:      urlArticle,
		})
	})

	s.logger.Printf("scrapping of medium %s finished successfully", mediumConfig.Name)
	return articles
}

func (s *Scrapper) getUrl(mediumUrl, articleUrl string) string {
	var fullUrl string
	if isUrlValid(articleUrl) {
		return articleUrl
	}

	if strings.HasSuffix(mediumUrl, "/") &&
		strings.HasPrefix(articleUrl, "/") {
		fullUrl = fmt.Sprintf("%s%s", mediumUrl, articleUrl[1:])
		if isUrlValid(fullUrl) {
			return fullUrl
		}
	}

	if !strings.HasSuffix(mediumUrl, "/") &&
		!strings.HasPrefix(articleUrl, "/") {
		fullUrl = fmt.Sprintf("%s/%s", mediumUrl, articleUrl)
		if isUrlValid(fullUrl) {
			return fullUrl
		}
	}

	fullUrl = fmt.Sprintf("%s%s", mediumUrl, articleUrl)
	if isUrlValid(fullUrl) {
		return fullUrl
	}

	urlWithoutDuplications := tryRemoveDuplicates(fullUrl)
	if urlWithoutDuplications == fullUrl || !isUrlValid(urlWithoutDuplications) {
		s.logger.Printf("failed to find the correct URL for the article. Medium URL (%s), article URL (%s)",
			mediumUrl, articleUrl)
		return ""
	}
	return urlWithoutDuplications

}

func isUrlValid(url string) bool {
	resp, err := http.Get(url)
	return err == nil && (resp.StatusCode >= 200 && resp.StatusCode <= 299)
}

func tryRemoveDuplicates(url string) string {
	var newURL string
	parts := deleteEmpty(strings.Split(url, "/"))

	newURL = parts[0]
	if strings.HasPrefix(newURL, "http") {
		newURL += "/"
	}

	for i, part := range parts {
		if i == 0 {
			continue
		}
		if part == parts[i-1] {
			continue
		}

		newURL = fmt.Sprintf("%s/%s", newURL, part)
	}

	return newURL
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
