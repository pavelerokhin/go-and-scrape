package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/pavelerokhin/go-and-scrape/business"
)

var wg sync.WaitGroup

func main() {
	mediums, err := business.ReadMediumConfig("medium-config.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, medium := range mediums.Mediums {
		wg.Add(1)
		m := medium.Medium
		go func() {
			err := business.ScrapeAndPersist(&m, &wg)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}
