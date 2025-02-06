package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/adnanibrahi0102/seo-analyzer-golang/models"
)

const apiURL = "https://api-inference.huggingface.co/models/facebook/bart-large-cnn"
const apiKey = ""

func PrepareDataForHuggingFaceModel(data models.RequestData) string {
	// Combine the title, description, keywords, headings, and paragraphs into a single content string
	var combinedContent []string

	//start with the title
	if data.Title != "" {
		combinedContent = append(combinedContent, data.Title)
	}

	// Add the description
	if data.Description != "" {
		combinedContent = append(combinedContent, data.Description)
	}

	// Add the keywords

	if data.Keywords != "" {
		combinedContent = append(combinedContent, data.Keywords)
	}

	// Add the headings

	if len(data.Headings) > 0 {
		combinedContent = append(combinedContent, "Headings: "+strings.Join(data.Headings, "\n"))
	}

	// Add the paragraphs
	if len(data.Paragraphs) > 0 {
		combinedContent = append(combinedContent, "Paragraphs: "+strings.Join(data.Paragraphs, "\n"))
	}

	// Add the links
	if len(data.URl) > 0 {
		combinedContent = append(combinedContent, "Links: "+strings.Join(data.URl, "\n"))
	}
	log.Printf("Combined content to send to Hugging Face API: %s", strings.Join(combinedContent, "\n"))
	// Join all parts into a single string
	return strings.Join(combinedContent, "\n")

}

func CallingHuggingFaceApi(content string) (string, error) {
	// Create a request payload that matches the expected Hugging Face API format
	reqData := map[string]string{
		"inputs": content, // Hugging Face expects `inputs` as the key
	}

	// Convert the request data to JSON
	reqJson, err := json.Marshal(reqData)
	if err != nil {
		log.Printf("Error marshalling request data: %v", err)
		return "", err
	}

	// Make the POST request to Hugging Face API
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqJson))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return "", err
	}

	// Log the request body to ensure it's properly formed
	log.Printf("Request body to Hugging Face API: %s", string(reqJson))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return "", err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return "", err
	}

	// Log the response body to see if the summary is returned
	log.Printf("Response body from Hugging Face API: %s", string(body))

	// Parse the response (expected to be an array of objects)
	var responseData []map[string]string
	err = json.Unmarshal(body, &responseData)
	if err != nil || len(responseData) == 0 {
		log.Printf("Error unmarshalling response data: %v", err)
		return "", err
	}

	// Extract the summary text
	summary, exists := responseData[0]["summary_text"]
	if !exists {
		log.Printf("Summary text not found in response")
		return "", nil
	}

	return summary, nil
}
