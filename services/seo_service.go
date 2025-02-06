package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/adnanibrahi0102/seo-analyzer-golang/config"
	"github.com/adnanibrahi0102/seo-analyzer-golang/models"
	"github.com/adnanibrahi0102/seo-analyzer-golang/repository"
	"github.com/gocolly/colly"
)

func AnalyzeSeo(url string) (*models.SEO_REPORT, error) {
	ctx := context.Background()

	// ✅ Check if data exists in Redis cache
	cachedData, err := config.RedisClient.Get(ctx, url).Result()
	if err == nil {
		var seoData models.SEO_REPORT
		err := json.Unmarshal([]byte(cachedData), &seoData)
		if err != nil {
			log.Println("Error unmarshaling cached data:", err)
			return nil, err
		}
		return &seoData, nil
	}

	// ✅ Initialize Colly collector
	seoData := &models.SEO_REPORT{URL: url}
	collector := colly.NewCollector()

	// ✅ Scrape title, description, and keywords
	collector.OnHTML("title", func(e *colly.HTMLElement) {
		seoData.Title = e.Text
	})
	collector.OnHTML(`meta[name="description"]`, func(e *colly.HTMLElement) {
		seoData.Description = e.Attr("content")
	})
	collector.OnHTML(`meta[name="keywords"]`, func(e *colly.HTMLElement) {
		seoData.Keywords = e.Attr("content")
	})

	// ✅ Scrape Headings (h1, h2, h3, etc.)
	var headings []string
	collector.OnHTML("h1, h2, h3, h4, h5, h6", func(e *colly.HTMLElement) {
		headings = append(headings, e.Text)
	})
	collector.OnScraped(func(r *colly.Response) {
		jsonBytes, _ := json.Marshal(headings)
		seoData.Headings = string(jsonBytes)
	})

	// ✅ Scrape Paragraphs (`p` tags)
	var paragraphs []string
	collector.OnHTML("p", func(e *colly.HTMLElement) {
		paragraphs = append(paragraphs, e.Text)
	})
	collector.OnScraped(func(r *colly.Response) {
		jsonBytes, _ := json.Marshal(paragraphs)
		seoData.Paragraphs = string(jsonBytes)
	})

	// ✅ Scrape All Links (`a` tags)
	var links []string
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		links = append(links, e.Attr("href"))
	})
	collector.OnScraped(func(r *colly.Response) {
		jsonBytes, _ := json.Marshal(links)
		seoData.Links = string(jsonBytes)
	})

	// ✅ Scrape Canonical URL (if available)
	collector.OnHTML(`link[rel="canonical"]`, func(e *colly.HTMLElement) {
		seoData.Canonical = e.Attr("href")
	})

	// ✅ Scrape OpenGraph & Twitter Metadata
	ogTags := make(map[string]string)
	twitterTags := make(map[string]string)
	collector.OnHTML(`meta[property^="og:"]`, func(e *colly.HTMLElement) {
		ogTags[e.Attr("property")] = e.Attr("content")
	})
	collector.OnHTML(`meta[name^="twitter:"]`, func(e *colly.HTMLElement) {
		twitterTags[e.Attr("name")] = e.Attr("content")
	})
	collector.OnScraped(func(r *colly.Response) {
		ogTagsJSON, _ := json.Marshal(ogTags)
		twitterTagsJSON, _ := json.Marshal(twitterTags)
		seoData.OgTags = string(ogTagsJSON)
		seoData.TwitterTags = string(twitterTagsJSON)
	})

	// ✅ Scrape img alt attributes
	collector.OnHTML(`img`, func(e *colly.HTMLElement) {
		seoData.ImageAlt += e.Attr("alt") + ", "
	})

	// ✅ Wait for scraping to complete
	collector.OnScraped(func(r *colly.Response) {
		// Save SEO report in Redis
		seoDataJSON, _ := json.Marshal(seoData)
		config.RedisClient.Set(ctx, url, seoDataJSON, 0)

		// Save SEO report in the database
		err := repository.SaveSeoReport(seoData)
		if err != nil {
			log.Println("Error saving to database:", err)
		}
	})

	// Start scraping
	err = collector.Visit(url)
	if err != nil {
		log.Println("Failed to visit URL:", err)
		return nil, err
	}

	return seoData, nil
}
