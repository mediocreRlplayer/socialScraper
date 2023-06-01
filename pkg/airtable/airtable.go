package airtable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Record struct {
	// backticks are used to specify the JSON key for the struct field
	Id string `json:"id"`
	Fields struct {
		Website string `json:"website"`
	} `json:"fields"`
}

type AirtableResponse struct {
	Records []Record `json:"records"`
}

func FetchWebsitesFromAirtable(apiKey, baseID, tableName string) ([]Record, error) {
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
	records := airtableResp.Records

	return records, nil
}

// AppendSocialLinksToTable appends the social links to the same table in the same row if possible,
// or appends them to a different table and copies over the name.
func AppendSocialLinksToTable(socialLinks []string, id,  websiteName, baseID, tableName, apiKey string) error {
	// check to make sure social links are not empty, if empty return out of function early 
	if len(socialLinks) == 0 {
		return nil
	}
	var twitter, instagram, facebook, linkedin string;
	for _, link := range socialLinks {
		switch {
		case strings.Contains(link, "twitter"):
			twitter = link
		case strings.Contains(link, "instagram"):
			instagram = link
		case strings.Contains(link, "facebook"):
			facebook = link
		default:
			linkedin = link
		}
	}
	// Create the JSON payload for the new record or record update
	data := map[string]interface{}{
		"records": []map[string]interface{}{
			{
			"id": id,
			"fields": map[string]interface{}{
				 "twitter": twitter,
				 "instagram": instagram,
				 "linkedin": linkedin,
				 "facebook": facebook,
		},
	},
},
}
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	payme := bytes.NewBuffer(payload)
	fmt.Printf("payload is %v\n", payme)

	// Create the HTTP request to append the social links
	url := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", baseID, tableName)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to append social links to Airtable. Status code: %d", resp.StatusCode)
	}

	return nil
}