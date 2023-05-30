package airtable

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Record struct {
	Fields struct {
		Website string `json:"website"`
	} `json:"fields"`
}

type AirtableResponse struct {
	Records []Record `json:"records"`
}

func FetchWebsitesFromAirtable(apiKey, baseID, tableName string) ([]string, error) {
	// Send a GET request to the Airtable API
	url := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", baseID, tableName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the API response
	var airtableResp AirtableResponse
	err = json.NewDecoder(resp.Body).Decode(&airtableResp)
	if err != nil {
		return nil, err
	}

	// Extract website URLs into a slice
	websites := make([]string, 0)
	for _, record := range airtableResp.Records {
		websites = append(websites, record.Fields.Website)
	}

	return websites, nil
}
