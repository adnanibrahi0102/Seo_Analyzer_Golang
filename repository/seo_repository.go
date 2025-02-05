package repository

import (
	"github.com/adnanibrahi0102/seo-analyzer-golang/config"
	"github.com/adnanibrahi0102/seo-analyzer-golang/models"
)

func SaveSeoReport(report *models.SEO_REPORT) error {
	// save the SEO report to the database
	return config.DB.Create(report).Error
}

func GetSeoReport(url string)(*models.SEO_REPORT, error){
	var report models.SEO_REPORT
    result := config.DB.Where("url = ?", url).First(&report)
    return &report, result.Error
}