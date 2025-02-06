package routes

import (
	"github.com/adnanibrahi0102/seo-analyzer-golang/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/analyze-seo", handlers.AnalyzeSeoHandler)
	// test route
	router.GET("/api/v1/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	router.POST("/api/v1/summarize-text", handlers.SummarizeTextHandler)
	return router
}