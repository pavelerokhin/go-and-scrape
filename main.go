package main

import (
	"fmt"
	"os"

	"github.com/pavelerokhin/go-and-scrape/business"
)

func main() {
	mediums, err := business.ReadMediumConfig("medium-config.yaml")
	checkErr(err)

	articles, err := business.ScrapMedium(&mediums.Mediums[0].Medium)
	checkErr(err)

	if len(articles) > 0 {
		err = business.WriteCSV(articles, &mediums.Mediums[0].Medium)
	}
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
