package storage

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

type CSVRepo struct {
	File   *os.File
	logger *log.Logger
	Writer *csv.Writer
}

// NewCSVRepo is a builder of CSV repository
func NewCSVRepo(logger *log.Logger, medium *configs.MediumConfig) (Storage, error) {
	file, err := os.Create(medium.FileName)
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(file)

	return &CSVRepo{File: file, logger: logger, Writer: writer}, nil
}

// GetArticleByID gets article with `id` from the CSV file
func (r *CSVRepo) GetArticleByID(id int) *entities.Article {

	return nil
}

// GetMediumByID gets medium with `id` from the CSV file
func (r *CSVRepo) GetMediumByID(id int) *entities.Medium {

	return nil
}

// GetMediumByURL gets medium with `id` from the CSV file
func (r *CSVRepo) GetMediumByURL(url string) *entities.Medium {

	return nil
}

// Save writes the scrapped article into the CSV file
func (r *CSVRepo) SaveArticle(a *entities.Article) *entities.Article {
	defer r.Writer.Flush()
	err := r.Writer.Write(a.ToSlice())
	if err != nil {
		r.logger.Println(err)
	}
	return a
}

func (r *CSVRepo) SaveMedium(article *entities.Medium) *entities.Medium {

	return nil
}

func (r *CSVRepo) writeHeaders(a *entities.Article) error {
	defer r.Writer.Flush()
	return r.Writer.Write(a.GetHeaders())
}
