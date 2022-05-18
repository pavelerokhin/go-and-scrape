package storage

import (
	"fmt"
	"github.com/pavelerokhin/go-and-scrape/scrapper/models/entities"
	"github.com/pavelerokhin/go-and-scrape/scrapper/models/nlp"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

type SQLiteRepo struct {
	DB     *gorm.DB
	logger *log.Logger
}

func NewSQLiteArticleRepo(dbFileName string, logger *log.Logger) (Storage, error) {
	if dbFileName == "" {
		return &SQLiteRepo{}, fmt.Errorf("database name is empty")
	}

	sql, err := gorm.Open(sqlite.Open(fmt.Sprintf("../database/%s.db", dbFileName)), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
	})
	if err != nil {
		return &SQLiteRepo{}, err
	}

	err = sql.AutoMigrate(&entities.Medium{})
	if err != nil {
		return &SQLiteRepo{}, err
	}

	err = sql.AutoMigrate(&entities.ArticlePreview{})
	if err != nil {
		return &SQLiteRepo{}, err
	}

	err = sql.AutoMigrate(&entities.Article{})
	if err != nil {
		return &SQLiteRepo{}, err
	}

	err = sql.AutoMigrate(&nlp.ArticlesLemmas{})
	if err != nil {
		return &SQLiteRepo{}, err
	}

	return &SQLiteRepo{DB: sql, logger: logger}, nil
}

// GetArticleByID gets article with `id` from the SQLite DB
func (r *SQLiteRepo) GetArticleByID(id int) *entities.ArticlePreview {
	r.logger.Printf("getting article with ID %d", id)
	var article *entities.ArticlePreview
	tx := r.DB.Where("id = ?", id).Find(&article)
	if tx.RowsAffected != 0 {
		return article
	}
	r.logger.Printf("article with ID %v not found", id)
	return nil
}

// GetArticleByID gets article with `url` (relative) from the SQLite DB
// it is supposed to be unique
func (r *SQLiteRepo) GetArticleByURL(url string) *entities.ArticlePreview {
	var article *entities.ArticlePreview
	tx := r.DB.Where("relative_url = ?", url).Find(&article)
	if tx.RowsAffected != 0 {
		return article
	}
	return nil
}

// GetMediumByID gets medium with `id` from the SQLite DB
func (r *SQLiteRepo) GetMediumByID(id int) *entities.Medium {
	r.logger.Printf("getting medium with ID %d", id)
	var medium *entities.Medium
	tx := r.DB.Where("id = ?", id).Find(&medium)
	if tx.RowsAffected != 0 {
		return medium
	}
	r.logger.Printf("medium with ID %v not found", id)
	return nil
}

// GetMediumByURL gets medium with `url` from the SQLite DB
func (r *SQLiteRepo) GetMediumByURL(url string) *entities.Medium {
	r.logger.Printf("checking in SQLite DB if medium with URL %s exist", url)
	var medium *entities.Medium
	tx := r.DB.Where("url = ?", url).Find(&medium)
	if tx.RowsAffected != 0 {
		r.logger.Printf("medium %s has been found", url)
		return medium
	}
	r.logger.Printf("medium with URL %s not found", url)
	return nil
}

func (r *SQLiteRepo) SaveArticle(a *entities.ArticlePreview) *entities.ArticlePreview {
	tx := r.DB.Create(&a)
	if tx.Error != nil {
		r.logger.Printf("error saving article %e", tx.Error)
		return nil
	}
	return a
}

// Save saves medium with all scrapped articles to the SQLite DB
func (r *SQLiteRepo) SaveMedium(m *entities.Medium) *entities.Medium {
	tx := r.DB.Create(&m)
	if tx.Error != nil {
		r.logger.Printf("error saving medium %e", tx.Error)
		return nil
	}
	return m
}
