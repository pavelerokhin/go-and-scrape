package business

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/pavelerokhin/go-and-scrape/models/configs"
)

// ReadMediumConfig reads `configs` file and returns list of Mediums
func ReadMediumConfig(config string) (*configs.ConfigFile, error) {
	b, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}

	mediums := &configs.ConfigFile{}
	err = yaml.Unmarshal(b, mediums)
	if err != nil {
		return nil, err
	}

	return mediums, nil
}
