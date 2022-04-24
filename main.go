package main

import (
	"fmt"
	"os"

	"github.com/pavelerokhin/go-and-scrape/business"
	"github.com/pavelerokhin/go-and-scrape/models"
)

func main() {
	meduza := models.Medium{
		URL: "https://meduza.io/", 
		CsvName: "meduza2.csv",
		HTMLTags: models.HTMLTags{
			Article: "article",
			Tag: ".RichBlock-tag",
			Title: ".BlockTitle-first",
			Subtitle: ".BlockTitle-second",
			URL: ".Link-root",
		},
	}

	articles, err := business.ScrapMedium(&meduza)
	checkErr(err)
	
	if len(articles) > 0 {
		err = business.WriteCSV(articles, &meduza)
	}
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

