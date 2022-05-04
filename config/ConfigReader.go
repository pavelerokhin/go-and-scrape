package config

import (
	"fmt"
	"io/ioutil"

	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"gopkg.in/yaml.v2"
)

// ReadConfig reads `configs` file and returns list of Mediums and settings for the persistence
func ReadConfig(configFilePath string) (configs.MediaConfig, *configs.PersistenceConfig, error) {
	mediumConfigSection, err := unmarshallPrototype(configFilePath, &configs.ScrapeSection{})
	if err != nil {
		return nil, nil, err
	}

	persistenceConfigSection, err := unmarshallPrototype(configFilePath, &configs.PersistenceSection{})
	if err != nil {
		return nil, nil, err
	}

	return mediumConfigSection.(*configs.ScrapeSection).Mediums,
		&persistenceConfigSection.(*configs.PersistenceSection).PersistenceConfig, nil
}

func unmarshallPrototype(configFilePath string, prototype interface{}) (output interface{}, err error) {
	bytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading the config file %e", err)
	}

	err = yaml.Unmarshal(bytes, prototype)
	if err != nil {
		return nil, fmt.Errorf("error reading the config file %e", err)
	}

	return prototype, nil
}
