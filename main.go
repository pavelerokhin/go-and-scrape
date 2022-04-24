package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"

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

	response, err := http.Get(meduza.URL)
	defer response.Body.Close()
	checkErr(err)

	if response.StatusCode >= 400 {
		fmt.Printf("Status code: %v", response.StatusCode)
		os.Exit(1)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	checkErr(err)

	newsContainer := document.Find(meduza.HTMLTags.Article)
	if newsContainer.Size() == 0 {
		fmt.Println("no news")
		os.Exit(0)
	}
	fmt.Printf("%d articles has been fond\n", newsContainer.Size())


	file, err := os.Create(meduza.CsvName)
	checkErr(err)

	var articles []models.Article
	
	newsContainer.Each(func(i int, item *goquery.Selection) {
		tag := strings.TrimSpace(item.Find(meduza.HTMLTags.Tag).Text())
		title := strings.TrimSpace(item.Find(meduza.HTMLTags.Title).Text())
		subtitle := strings.TrimSpace(item.Find(meduza.HTMLTags.Subtitle).Text())
		urlArticle, _ := item.Find(meduza.HTMLTags.URL).Attr("href")
		urlArticle = fmt.Sprintf("%s%s", meduza.URL, urlArticle)
		
		articles = append(articles, models.Article{
			Tag: tag,
			Title: title,
			Subtitle: subtitle,
			URL: urlArticle,
		})
	})

	if len(articles) > 0 {
		writer := csv.NewWriter(file)
		defer writer.Flush()
		writer.Write(articles[0].GetHeaders())
		for i:=0; i<len(articles); i++ {
			writer.Write(articles[i].ToSlice())
		}
		
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func writeFile(data, filename string) {
	file, err := os.Create(filename)
	defer file.Close()
	checkErr(err)

	file.WriteString(data)
}