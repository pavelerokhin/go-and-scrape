package storage

import (
	"fmt"
	"github.com/pavelerokhin/go-and-scrape/models/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

type SQLiteArticleRepo struct {
	DB *gorm.DB
}

func NewSQLiteArticleRepo(dbFileName string) (*SQLiteArticleRepo, error) {
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

	return &SQLiteArticleRepo{DB: sql}, nil
}

// GetArticle gets article with `id` from the SQLite DB
func (r *SQLiteArticleRepo) GetArticle(id int) (*entities.Article, error) {
	var article *entities.Article
	tx := r.DB.Where("id = ?", id).Find(&article)

	if tx.RowsAffected != 0 {
		return article, nil
	}

	return nil, fmt.Errorf("article with ID %v not found", id)
}

// GetMedium gets medium with `id` from the SQLite DB
func (r *SQLiteArticleRepo) GetMedium(id int) (*entities.Medium, error) {
	var medium *entities.Medium
	tx := r.DB.Where("id = ?", id).Find(&medium)

	if tx.RowsAffected != 0 {
		return medium, nil
	}

	return nil, fmt.Errorf("medium with ID %v not found", id)
}

// Save saves medium with all scrapped articles to the SQLite DB
func (r *SQLiteArticleRepo) Save(m *entities.Medium) (*entities.Medium, error) {
	tx := r.DB.Create(&m)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return m, nil
}
