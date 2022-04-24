package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://meduza.io/"

	response, err := http.Get(url)
	defer response.Body.Close()
	checkErr(err)

	if response.StatusCode >= 400 {
		fmt.Printf("Status code: %v", response.StatusCode)
		os.Exit(1)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	checkErr(err)

	newsContainer := document.Find("article")
	if newsContainer.Size() == 0 {
		fmt.Println("no news")
		os.Exit(0)
	}
	fmt.Printf("%d articles has been fond\n", newsContainer.Size())


	file, err := os.Create("meduza.csv")
	checkErr(err)

	writer := csv.NewWriter(file)
	newsContainer.Each(func(i int, item *goquery.Selection) {
		tag := strings.TrimSpace(item.Find(".RichBlock-tag").Text())
		title := strings.TrimSpace(item.Find(".BlockTitle-first").Text())
		subtitle := strings.TrimSpace(item.Find(".BlockTitle-second").Text())
		urlArticle, _ := item.Find(".Link-root").Attr("href")
		urlArticle = fmt.Sprintf("%s%s", url, urlArticle)
		
		fmt.Printf("%s --- %s --- %s\n%s\n\n", tag, title, subtitle, urlArticle)
		writer.Write([]string{strconv.Itoa(i), tag, title, subtitle, urlArticle})
	})

	writer.Flush()
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