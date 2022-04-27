package store

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"

	"github.com/pavelerokhin/go-and-scrape/models"
)

type SQLiteArticleRepo struct {
	DB *gorm.DB
}

func NewSQLiteArtcleRepo(medium *models.Medium) (*SQLiteArticleRepo, error) {
	if medium.FileName == "" {
		return nil, fmt.Errorf("database name is empty")
	}

	sql, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", medium.FileName)), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
	})
	if err != nil {
		return nil, err
	}

	err = sql.AutoMigrate(&models.Article{})
	if err != nil {
		return nil, err
	}

	return &SQLiteArticleRepo{DB: sql}, nil
}

func (r *SQLiteArticleRepo) Get(id int) (*models.Article, error) {
	var article *models.Article
	tx := r.DB.Where("id = ?", id).Find(&article)

	if tx.RowsAffected != 0 {
		return article, nil
	}

	return nil, fmt.Errorf("article with ID %v not found", id)
}

func (r *SQLiteArticleRepo) Save(a *models.Article) (*models.Article, error) {
	tx := r.DB.Create(&a)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return a, nil
}
