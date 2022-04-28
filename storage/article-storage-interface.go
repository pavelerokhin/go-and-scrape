package storage

import (
	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

type ArticleStorage interface {
	GetArticle(id int) (*entities.Article, error)
	GetMedium(id int) (*entities.Medium, error)
	Save(article *entities.Medium) (*entities.Medium, error)
}
