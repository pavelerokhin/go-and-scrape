package models

type ConfigFile struct {
	Mediums Mediums `yaml:"mediums"`
}

type Mediums []struct {
	Medium Medium `yaml:"medium"`
}

// Medium is a basic news portal model
type Medium struct {
	Name     string   `yaml:"name"`
	URL      string   `yaml:"medium-url"`
	FileName string   `yaml:"file-name"`
	HTMLTags HTMLTags `yaml:"html-tags"`
}

// HTMLTags medium specific set of CSS-selectors for
// the requested parts of the article to scrape
type HTMLTags struct {
	Article  string `yaml:"article"`
	Tag      string `yaml:"tag"`
	Title    string `yaml:"title"`
	Subtitle string `yaml:"subtitle"`
	URL      string `yaml:"article-url"`
}
