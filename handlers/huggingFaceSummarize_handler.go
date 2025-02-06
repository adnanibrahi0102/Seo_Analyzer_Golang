package handlers

import (
	"net/http"

	"github.com/adnanibrahi0102/seo-analyzer-golang/models"
	"github.com/adnanibrahi0102/seo-analyzer-golang/services"
	"github.com/gin-gonic/gin"
)

func SummarizeTextHandler(c *gin.Context) {
	// Bind the incoming request JSON body to a struct
	var requestData models.RequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the content to summarize (title, description, keywords, headings, paragraphs, links)
	contentToSummarize := services.PrepareDataForHuggingFaceModel(requestData)

	// Call the Hugging Face API to summarize the content

	summary, err := services.CallingHuggingFaceApi(contentToSummarize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the summary in the response
	c.JSON(200, gin.H{"summary": summary})

}
