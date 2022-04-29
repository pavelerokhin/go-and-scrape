package modules

import (
	"fmt"
	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ScrapMedium is a function that scpapa medium and returns a slice of Article
// objects and error
func ScrapMedium(mediumConfig *configs.MediumConfig) ([]entities.Article, error) {
	response, err := http.Get(mediumConfig.URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("status code: %v", response.StatusCode)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	newsContainer := document.Find(mediumConfig.HTMLTags.Article)
	if newsContainer.Size() == 0 {
		return nil, fmt.Errorf("no news")
	}

	fmt.Printf("%d articles has been found for the medium %s\n", newsContainer.Size(),
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

	return articles, nil
}
