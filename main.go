package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func scrapeSocialLinks(url string) ([]string, error) {
	var socialLinks []string

	// Send a GET request to the website URL
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check if the response was successful (status code 200)
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch the website. Status code: %d", response.StatusCode)
	}

	// Parse the HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	// Find the HTML elements containing social links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && isSocialLink(href) {
			socialLinks = append(socialLinks, href)
		}
	})

	return socialLinks, nil
}

// Helper function to check if a URL is a social link
func isSocialLink(url string) bool {
	// Convert the URL to lowercase for case-insensitive comparison
	lowercaseURL := strings.ToLower(url)

	// List of social media platforms
	socialPlatforms := []string{"instagram", "twitter", "facebook", "linkedin"}

	// Check if the lowercaseURL contains any of the social platforms
	for _, platform := range socialPlatforms {
		if strings.Contains(lowercaseURL, platform) {
			return true
		}
	}

	return false
}

func main() {
	url := "https://galaxy.com"

	socialLinks, err := scrapeSocialLinks(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Social Links:")
	for _, link := range socialLinks {
		fmt.Println(link)
	}
}
