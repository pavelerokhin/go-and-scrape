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
				Article:     article.Article,
			})
	}

	return normalizedArticles
}

func addWhitespaceAroundPeriod(s string) string {
	period := regexp.MustCompile(`[a-zA-z](\.)[a-zA-z]`)
	return period.ReplaceAllLiteralString(s, " . ")
}

func normalizeText(s string) string {
	return addWhitespaceAroundPeriod(normalizeWhitespaces(stripPunctuation(s)))
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
