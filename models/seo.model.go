package models

import "gorm.io/gorm"

type SEO_REPORT struct {
	gorm.Model
	URL         string `json:"url" gorm:"unique"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Headings    string `json:"headings"` // Store all h1, h2, h3 as JSON
	Canonical   string `json:"canonical"`
	OgTags      string `json:"og_tags"`      // Store OpenGraph data as JSON
	TwitterTags string `json:"twitter_tags"` // Store Twitter metadata as JSON
	ImageAlt    string `json:"imageAlt"`     // Store `alt` text of images
	Links       string `json:"links"`        // Store all extracted links as JSON
	Paragraphs  string `json:"paragraphs"`   // Store all paragraphs as JSON
}

type RequestData struct {
    Title string `json:"title"`
    Description string `json:"description"`
    Keywords string `json:"keywords"`
    Headings []string `json:"headings"`
    Paragraphs []string `json:"paragraphs"`
    URl []string `json:"links"`
}

type ResponseData struct {
	SummaryText string `json:"summary_text"`
}