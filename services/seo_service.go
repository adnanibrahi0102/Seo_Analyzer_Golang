package services

import (
	"github.com/adnanibrahi0102/seo-analyzer-golang/config"
	"github.com/adnanibrahi0102/seo-analyzer-golang/models"
	"github.com/adnanibrahi0102/seo-analyzer-golang/repository"
	"github.com/gocolly/colly"
)

func AnalyzeSeo(url string) (*models.SEO_REPORT, error) {
	// check redis cache for the SEO report
	// if found, return the report
	// if not found, analyze the website and save the report to the database
	ctx := config.Ctx 
	cachedData, err := config.RedisClient.Get(ctx, url).Result()

	if err == nil {
		return &models.SEO_REPORT{URL: url, Title: cachedData}, nil
	}

	// scrape the website
	seoData := &models.SEO_REPORT{URL: url}
	collector := colly.NewCollector()

	collector.OnHTML("title", func(e *colly.HTMLElement) {
		seoData.Title = e.Text
	})

	collector.OnHTML("meta[name='description']", func(e *colly.HTMLElement) {
		seoData.Description = e.Attr("content")
	})
	
	collector.Visit(url)

	// save the seo report in the redis
	config.RedisClient.Set(ctx, url, seoData.Title, 0)

	// save the seo report in the database
	err = repository.SaveSeoReport(seoData)
	if err != nil {
		return nil, err
	}

	return seoData, nil

}
