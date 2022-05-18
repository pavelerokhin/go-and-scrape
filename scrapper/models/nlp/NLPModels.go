package nlp

import "github.com/pavelerokhin/go-and-scrape/scrapper/models/entities"

// ArticlesLemmas is a basic ner metrics persistence
type ArticlesLemmas struct {
	ArticleID uint `gorm:"article_id"`
	Article   entities.Article
	Lemma     string `gorm:"lemma"`
	Type      string `gorm:"type"`
	Count     string `gorm:"count"`
}
