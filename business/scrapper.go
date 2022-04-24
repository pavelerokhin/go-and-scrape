package business

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pavelerokhin/go-and-scrape/models"
)

func ScrapMedium(medium *models.Medium) ([]models.Article, error) {
	response, err := http.Get(medium.URL)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("Status code: %v", response.StatusCode)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	newsContainer := document.Find(medium.HTMLTags.Article)
	if newsContainer.Size() == 0 {
		return nil, fmt.Errorf("no news")
	}
	
	fmt.Printf("%d articles has been fond\n", newsContainer.Size())
	var articles []models.Article
	newsContainer.Each(func(i int, item *goquery.Selection) {
		tag := strings.TrimSpace(item.Find(medium.HTMLTags.Tag).Text())
		title := strings.TrimSpace(item.Find(medium.HTMLTags.Title).Text())
		subtitle := strings.TrimSpace(item.Find(medium.HTMLTags.Subtitle).Text())
		urlArticle, _ := item.Find(medium.HTMLTags.URL).Attr("href")
		urlArticle = fmt.Sprintf("%s%s", medium.URL, urlArticle)
		
		articles = append(articles, models.Article{
			Tag: tag,
			Title: title,
			Subtitle: subtitle,
			URL: urlArticle,
		})
	})

	return articles, nil
}
