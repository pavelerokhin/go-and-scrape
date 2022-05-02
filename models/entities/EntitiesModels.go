package entities

import (
	"fmt"
	"gorm.io/gorm"
)

// Medium is a basic news portal model
type Medium struct {
	gorm.Model
	Name            string `gorm:"name"`
	URL             string `gorm:"medium_url"`
	ArticlePreviews []ArticlePreview
}

func (m *Medium) ToString() string {
	return fmt.Sprintf("medium %s with ID %v, URL %s", m.Name, m.ID, m.URL)
}

// ArticlePreview is a basic news article model
type ArticlePreview struct {
	gorm.Model
	Tag      string `gorm:"tag"`
	Title    string `gorm:"title"`
	Subtitle string `gorm:"subtitle"`
	URL      string `gorm:"url"`
	MediumID uint
	Article  Article
}

// GetHeaders is a util function which gets CSV headers
func (a *ArticlePreview) GetHeaders() []string {
	return []string{"tag", "title", "subtitle", "url"}
}

// ToSlice is an util function which transforms struct ArticlePreview into string slice
func (a *ArticlePreview) ToSlice() []string {
	return []string{a.Tag, a.Title, a.Subtitle, a.URL}
}

// ToString is an util function which transforms struct ArticlePreview into a string (for better logging)
func (a *ArticlePreview) ToString() string {
	if a.Title != "" {
		return fmt.Sprintf("article ID %v of medium %v: %s", a.ID, a.MediumID, a.Title)
	}
	return fmt.Sprintf("article preview ID %v by medium %v", a.ID, a.MediumID)
}

// ArticlePreview is a basic news article model
type Article struct {
	gorm.Model
	Author           string `gorm:"author"`
	Date             string `gorm:"date"`
	Text             string `gorm:"text"`
	ArticlePreviewId uint
}
