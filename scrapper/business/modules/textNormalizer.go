package modules

import (
	"regexp"

	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

func Normalize(articles []entities.ArticlePreview) []entities.ArticlePreview {
	var normalizedArticles []entities.ArticlePreview
	for _, article := range articles {
		normalizedArticles = append(normalizedArticles,
			entities.ArticlePreview{
				Tag:         normalizeText(article.Tag),
				Title:       normalizeText(article.Title),
				Subtitle:    normalizeText(article.Subtitle),
				URL:         article.URL,
				RelativeURL: article.RelativeURL,
				MediumID:    article.MediumID,
				Article: entities.Article{
					Author: article.Article.Author,
					Date:   article.Article.Date,
					Text:   normalizeText(article.Article.Text),
				},
			})
	}

	return normalizedArticles
}

func addWhitespaceAroundPunctuation(s string) string {
	punctuation := regexp.MustCompile(`([^a-zA-Z])(\.|\,)([^a-zA-Z])`)
	return punctuation.ReplaceAllString(s, "$1$2 $3")
}

func normalizeText(s string) string {
	return addWhitespaceAroundPunctuation(normalizeWhitespaces(s))
}

func normalizeWhitespaces(s string) string {
	ws := regexp.MustCompile(`[ \s]`)
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllLiteralString(ws.ReplaceAllLiteralString(s, " "), " ")
}

func stripPunctuation(s string) string {
	r := regexp.MustCompile(`[()\[\].,\-"':;«»—!?]`)
	return r.ReplaceAllLiteralString(s, "")
}
