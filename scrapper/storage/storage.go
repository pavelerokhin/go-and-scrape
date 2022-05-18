package storage

import (
	"github.com/pavelerokhin/go-and-scrape/scrapper/models/entities"
)

type Storage interface {
	GetArticleByID(id int) *entities.ArticlePreview
	GetArticleByURL(url string) *entities.ArticlePreview
	GetMediumByID(id int) *entities.Medium
	GetMediumByURL(url string) *entities.Medium
	SaveMedium(article *entities.Medium) *entities.Medium
	SaveArticle(articles *entities.ArticlePreview) *entities.ArticlePreview
}
