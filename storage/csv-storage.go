package storage

import (
	"encoding/csv"
	"os"

	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

type CSVArticleRepo struct {
	File   *os.File
	Writer *csv.Writer
}

// NewCSVArticleRepo is a builder of CSV repository
func NewCSVArticleRepo(medium *configs.MediumConfig) (*CSVArticleRepo, error) {
	file, err := os.Create(medium.FileName)
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(file)

	return &CSVArticleRepo{File: file, Writer: writer}, nil
}

// GetArticleByID gets article with `id` from the CSV file
func (r *CSVArticleRepo) GetArticleByID(id int) (*entities.Article, error) {

	return nil, nil
}

// GetMediumByID gets medium with `id` from the CSV file
func (r *CSVArticleRepo) GetMediumByID(id int) (*entities.Article, error) {

	return nil, nil
}

// GetMediumByURL gets medium with `id` from the CSV file
func (r *CSVArticleRepo) GetMediumByURL(url string) (*entities.Article, error) {

	return nil, nil
}

// Save writes the scrapped article into the CSV file
func (r *CSVArticleRepo) Save(a *entities.Article) (*entities.Article, error) {
	defer r.Writer.Flush()
	err := r.Writer.Write(a.ToSlice())
	return a, err
}

func (r *CSVArticleRepo) WriteHeaders(a *entities.Article) error {
	defer r.Writer.Flush()
	return r.Writer.Write(a.GetHeaders())
}
