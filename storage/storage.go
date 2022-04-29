package storage

import (
	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

type Storage interface {
	GetArticleByID(id int) *entities.Article
	GetMediumByID(id int) *entities.Medium
	GetMediumByURL(url string) *entities.Medium
	SaveMedium(article *entities.Medium) *entities.Medium
	SaveArticle(articles *entities.Article) *entities.Article
}
