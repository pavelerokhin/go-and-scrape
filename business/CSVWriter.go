package business

import (
	"encoding/csv"
	"os"

	"github.com/pavelerokhin/go-and-scrape/models"
)

// WriteCSV writes the scrapped articles into the CSV file
func WriteCSV(articles []models.Article, medium *models.Medium) error {
	file, err := os.Create(medium.CsvName)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()
	
	writer.Write(articles[0].GetHeaders())
	for i:=0; i<len(articles); i++ {
		writer.Write(articles[i].ToSlice())
	}

	return nil
}
