package models

import "gorm.io/gorm"


type SEO_REPORT struct {
	gorm.Model  // This embeds GORM's model to provide fields like ID, CreatedAt, etc.
    URL         string `json:"url" gorm:"unique"`   // URL of the website
    Title       string `json:"title"`               // Title tag of the website
    Description string `json:"description"`         // Description meta tag of the website
}