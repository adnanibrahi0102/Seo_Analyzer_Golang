package handlers

import (
	"net/http"

	"github.com/adnanibrahi0102/seo-analyzer-golang/services"
	"github.com/gin-gonic/gin"
)

func AnalyzeSeoHandler(c *gin.Context) {
	url := c.Query("url")

	if url == "" {
		c.JSON(400, gin.H{"error": "url query parameter is required"})
		return
	}

	report, err := services.AnalyzeSeo(url)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": report})
}
