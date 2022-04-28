package business

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"github.com/pavelerokhin/go-and-scrape/storage"
)

func ScrapeAndPersist(storage *storage.SQLiteRepo, mediumConfig *configs.MediumConfig, wg *sync.WaitGroup) error {
	defer wg.Done()

	articles, err := ScrapMedium(mediumConfig)
	if err != nil {
		return err
	}

	var medium *entities.Medium
	medium, err = storage.GetMediumByURL(mediumConfig.URL)
	if err != nil {
		return err
	}

	if len(articles) > 0 {
		articles = normalizeArticlesNLP(articles)

		if medium.URL == "" {
			_, err = storage.SaveMedium(&entities.Medium{
				Name:     mediumConfig.Name,
				URL:      mediumConfig.URL,
				Articles: articles,
			})
		} else {
			for _, article := range articles {
				a := article
				a.MediumID = medium.ID
				_, err = storage.SaveArticle(&a)
			}
		}

	}
	if err != nil {
		return err
	}

	return nil
}

func normalizeArticlesNLP(articles []entities.Article) []entities.Article {
	var normalizedArtiles []entities.Article
	for _, article := range articles {
		normalizedArtiles = append(normalizedArtiles, entities.Article{
			Tag:      nlpManagerSmall(article.Tag),
			Title:    nlpManagerSmall(article.Title),
			Subtitle: nlpManagerSmall(article.Subtitle),
			URL:      article.URL,
			MediumID: article.MediumID,
		})
	}

	return normalizedArtiles
}

func nlpManagerBig(s string) []string {
	noPunctuation := stripPunctuation(s)
	words := splitWords(noPunctuation)

	var stems []string
	for _, word := range words {
		stems = append(stems, strings.ToLower(stem(word)))
	}

	wordsCountDict := countWords(stems)
	wordsRanked := rankByWordCount(wordsCountDict)
	fmt.Println(wordsRanked)
	return nil
}

func nlpManagerSmall(s string) string {
	return strings.ToLower(stripPunctuation(s))
}
