package storage

import (
	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

type ArticleStorage interface {
	GetArticleByID(id int) (*entities.Article, error)
	GetMediumByID(id int) (*entities.Medium, error)
	GetMediumByURL(url string) (*entities.Medium, error)
	SaveMedium(article *entities.Medium) (*entities.Medium, error)
	SaveArticle(articles *entities.Article) (*entities.Article, error)
}
