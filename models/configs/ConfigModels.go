package configs

type ConfigFile struct {
	Mediums MediumConfigs `yaml:"mediums"`
}

type MediumConfigs []struct {
	MediumConfig MediumConfig `yaml:"medium"`
}

// MediumConfig contains info for parsing configuration YAML for mediums
type MediumConfig struct {
	Name     string   `yaml:"name"`
	URL      string   `yaml:"medium-url"`
	FileName string   `yaml:"file-name"`
	HTMLTags HTMLTags `yaml:"html-tags"`
}

// HTMLTags medium specific set of CSS-selectors for
// the requested parts of the article to scrape
type HTMLTags struct {
	ID       uint
	Article  string `yaml:"article"`
	Tag      string `yaml:"tag"`
	Title    string `yaml:"title"`
	Subtitle string `yaml:"subtitle"`
	URL      string `yaml:"article-url"`
}