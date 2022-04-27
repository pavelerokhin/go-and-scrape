package business

import (
	"github.com/pavelerokhin/go-and-scrape/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// ReadMediumConfig reads `config` file and returns list of Mediums
func ReadMediumConfig(config string) (*models.ConfigFile, error) {
	b, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}

	mediums := &models.ConfigFile{}
	err = yaml.Unmarshal(b, mediums)
	if err != nil {
		return nil, err
	}

	return mediums, nil
}
