package storage

import (
	"fmt"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

type SQLiteRepo struct {
	DB     *gorm.DB
	logger *log.Logger
}

func NewSQLiteArticleRepo(dbFileName string, logger *log.Logger) (*SQLiteRepo, error) {
	if dbFileName == "" {
		return nil, fmt.Errorf("database name is empty")
	}

	sql, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", dbFileName)), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
	})
	if err != nil {
		return nil, err
	}

	err = sql.AutoMigrate(&entities.Medium{})
	if err != nil {
		return nil, err
	}

	err = sql.AutoMigrate(&entities.Article{})
	if err != nil {
		return nil, err
	}

	return &SQLiteRepo{DB: sql, logger: logger}, nil
}

// GetArticleByID gets article with `id` from the SQLite DB
func (r *SQLiteRepo) GetArticleByID(id int) *entities.Article {
	r.logger.Printf("getting article with ID %d", id)
	var article *entities.Article
	tx := r.DB.Where("id = ?", id).Find(&article)
	if tx.RowsAffected != 0 {
		return article
	}
	r.logger.Printf("article with ID %v not found", id)
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
	r.logger.Printf("getting medium with URL %s", url)
	var medium *entities.Medium
	tx := r.DB.Where("url = ?", url).Find(&medium)
	if tx.RowsAffected != 0 {
		return medium
	}
	r.logger.Printf("medium with URL %s not found", url)
	return nil
}

func (r *SQLiteRepo) SaveArticle(a *entities.Article) *entities.Article {
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
