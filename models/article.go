package models

// Article is a basic news article model
type Article struct {
	Tag string `json:"tag"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
	URL string `json:"url"`
}

// GetHeaders is a util function which gets CSV headers
func (a *Article) GetHeaders() []string {
	return []string{"tag", "title", "subtitle", "url"}
}

// ToSlice is a util function which transforms struct Article into string slice
func (a *Article) ToSlice() []string {
	return []string{a.Tag, a.Title, a.Subtitle, a.URL}
}