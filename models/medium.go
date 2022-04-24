package models

// Medium is a basic news portal model
type Medium struct {
	URL string
	CsvName string
	HTMLTags HTMLTags
}

// HTMLTags medium specific set of CSS-selectors for 
// the requested parts of the article to scrape
type HTMLTags struct {
	Article string
	Tag string
	Title string
	Subtitle string
	URL string
}