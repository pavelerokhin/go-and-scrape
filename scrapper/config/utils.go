package config

import (
	"fmt"
	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"os"
)

func CheckMediaConfig(mediaConfig *configs.MediaConfig) {
	if len(*mediaConfig) == 0 {
		fmt.Println("no media set")
		os.Exit(0)
	}
}
