package store

import "github.com/pavelerokhin/go-and-scrape/models"

type ArticleStorage interface {
	Get(id int) (*models.Article, error)
	Save(article *models.Article) (*models.Article, error)
}
