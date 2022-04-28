package storage

import (
	"fmt"
	"github.com/pavelerokhin/go-and-scrape/models/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

type SQLiteRepo struct {
	DB *gorm.DB
}

func NewSQLiteArticleRepo(dbFileName string) (*SQLiteRepo, error) {
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

	return &SQLiteRepo{DB: sql}, nil
}

// GetArticleByID gets article with `id` from the SQLite DB
func (r *SQLiteRepo) GetArticleByID(id int) (*entities.Article, error) {
	var article *entities.Article
	tx := r.DB.Where("id = ?", id).Find(&article)

	if tx.RowsAffected != 0 {
		return article, nil
	}

	return nil, fmt.Errorf("article with ID %v not found", id)
}

// GetMediumByID gets medium with `id` from the SQLite DB
func (r *SQLiteRepo) GetMediumByID(id int) (*entities.Medium, error) {
	var medium *entities.Medium
	tx := r.DB.Where("id = ?", id).Find(&medium)

	if tx.RowsAffected != 0 {
		return medium, nil
	}

	return nil, fmt.Errorf("medium with ID %v not found", id)
}

// GetMediumByURL gets medium with `url` from the SQLite DB
func (r *SQLiteRepo) GetMediumByURL(url string) (*entities.Medium, error) {
	var medium *entities.Medium
	r.DB.Where("url = ?", url).Find(&medium)

	return medium, nil
}

func (r *SQLiteRepo) SaveArticle(a *entities.Article) (*entities.Article, error) {
	tx := r.DB.Create(&a)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return a, nil
}

// Save saves medium with all scrapped articles to the SQLite DB
func (r *SQLiteRepo) SaveMedium(m *entities.Medium) (*entities.Medium, error) {
	tx := r.DB.Create(&m)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return m, nil
}
