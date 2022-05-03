package modules

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Scrapper struct {
	logger *log.Logger
}

func NewScrapper(logger *log.Logger) *Scrapper {
	return &Scrapper{logger: logger}
}

// Scrap is a function that scraps medium and returns a slice of ArticlePreview
// objects and error
func (s *Scrapper) Scrap(mediumConfig *configs.MediumConfig) []entities.ArticlePreview {
	s.logger.Printf("start to scrap medium %s", mediumConfig.Name)
	document, err := getDocument(mediumConfig.URL)
	if err != nil {
		s.logger.Printf("error sending the request: %e", err)
		return nil
	}

	newsContainer := document.Find(mediumConfig.HTMLArticlePreviewTags.Article)
	if newsContainer.Size() == 0 {
		s.logger.Println("no news found")
		return nil
	}

	s.logger.Printf("%d articles has been found for the medium %s", newsContainer.Size(),
		mediumConfig.Name)

	var articlePreviews []entities.ArticlePreview
	var c = make(chan entities.ArticlePreview, 2)
	newsContainer.Each(func(i int, item *goquery.Selection) {
		go articleScrapWorker(c, item, mediumConfig, s.logger)
	})

	for i := 0; i < newsContainer.Length(); i++ {
		articlePreviews = append(articlePreviews, <-c)
	}

	s.logger.Printf("scrapping of medium %s finished successfully", mediumConfig.Name)
	return articlePreviews
}

func articleScrapWorker(c chan entities.ArticlePreview, item *goquery.Selection, mediumConfig *configs.MediumConfig,
	logger *log.Logger) {

	var err error
	tag := strings.TrimSpace(item.Find(mediumConfig.HTMLArticlePreviewTags.Tag).Text())
	title := strings.TrimSpace(item.Find(mediumConfig.HTMLArticlePreviewTags.Title).Text())
	subtitle := strings.TrimSpace(item.Find(mediumConfig.HTMLArticlePreviewTags.Subtitle).Text())

	urlArticle, _ := item.Find(mediumConfig.HTMLArticlePreviewTags.URL).Attr("href")
	urlArticle, err = getUrl(mediumConfig.URL, urlArticle)

	article, err := getArticle(mediumConfig, urlArticle)
	if err != nil {
		logger.Printf("cannot parse article with URL %s for medium %s",
			urlArticle,
			mediumConfig.Name)
		return
	}
	c <- entities.ArticlePreview{
		Tag:      tag,
		Title:    title,
		Subtitle: subtitle,
		URL:      urlArticle,
		Article:  article,
	}
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

func getArticle(mediumConfig *configs.MediumConfig, url string) (entities.Article, error) {
	document, err := getDocument(url)
	if err != nil {
		return entities.Article{}, err
	}

	var author, date, text string
	author = document.Find(mediumConfig.HTMLArticleTags.Author).First().Text()
	date = document.Find(mediumConfig.HTMLArticleTags.Date).First().Text()
	text = document.Find(mediumConfig.HTMLArticleTags.Text).First().Text()

	return entities.Article{
		Author: author,
		Date:   date,
		Text:   text,
	}, nil

}

func getDocument(url string) (*goquery.Document, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("status code: %v", response.StatusCode)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading from the document: %e", err)
	}

	return document, nil
}

func getUrl(mediumUrl, articleUrl string) (string, error) {
	var fullUrl string

	if strings.HasSuffix(mediumUrl, "/") &&
		strings.HasPrefix(articleUrl, "/") {
		fullUrl = fmt.Sprintf("%s%s", mediumUrl, articleUrl[1:])
		if isUrlValid(fullUrl) {
			return fullUrl, nil
		}
	}

	if isUrlValid(articleUrl) {
		return articleUrl, nil
	}

	if !strings.HasSuffix(mediumUrl, "/") &&
		!strings.HasPrefix(articleUrl, "/") {
		fullUrl = fmt.Sprintf("%s/%s", mediumUrl, articleUrl)
		if isUrlValid(fullUrl) {
			return fullUrl, nil
		}
	}

	fullUrl = fmt.Sprintf("%s%s", mediumUrl, articleUrl)
	if isUrlValid(fullUrl) {
		return fullUrl, nil
	}

	urlWithoutDuplications := tryRemoveDuplicates(fullUrl)
	if urlWithoutDuplications == fullUrl || !isUrlValid(urlWithoutDuplications) {
		return "", fmt.Errorf("failed to find the correct URL for the article. Medium URL (%s), article URL (%s)",
			mediumUrl, articleUrl)
	}

	return urlWithoutDuplications, nil
}

func isUrlValid(url string) bool {
	resp, err := http.Get(url)
	return err == nil && (resp.StatusCode >= 200 && resp.StatusCode <= 299)
}

var numRegex = regexp.MustCompile(`\d+`)

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
		if part == parts[i-1] && !numRegex.MatchString(part) {
			continue
		}
		newURL = fmt.Sprintf("%s/%s", newURL, part)
	}

	return newURL
}
