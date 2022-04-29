package config

import (
	"fmt"
	"io/ioutil"

	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"gopkg.in/yaml.v2"
)

// ReadMediumConfig reads `configs` file and returns list of Mediums
func ReadMediumConfig(configFilePath string) (*configs.ConfigFile, error) {
	bytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading the config file %e", err)
	}

	configFile := &configs.ConfigFile{}
	err = yaml.Unmarshal(bytes, configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading the config file %e", err)
	}

	return configFile, nil
}
