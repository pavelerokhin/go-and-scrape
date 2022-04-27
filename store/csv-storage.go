package store

import (
	"encoding/csv"
	"os"

	"github.com/pavelerokhin/go-and-scrape/models"
)

type CSVArticleRepo struct {
	File   *os.File
	Writer *csv.Writer
}

// NewCSVArticleRepo is a builder of CSV repository
func NewCSVArticleRepo(medium *models.Medium) (*CSVArticleRepo, error) {
	file, err := os.Create(medium.FileName)
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(file)

	return &CSVArticleRepo{File: file, Writer: writer}, nil
}

// Get and article with `id` from the CSV file
func (r *CSVArticleRepo) Get(id int) (*models.Article, error) {

	return nil, nil
}

// Save writes the scrapped article into the CSV file
func (r *CSVArticleRepo) Save(a *models.Article) (*models.Article, error) {
	defer r.Writer.Flush()
	err := r.Writer.Write(a.ToSlice())
	return a, err
}

func (r *CSVArticleRepo) WriteHeaders(a *models.Article) error {
	defer r.Writer.Flush()
	return r.Writer.Write(a.GetHeaders())
}
