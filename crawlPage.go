package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Check if the current URL is on the same domain as the base URL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Error parsing base URL:", err)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Error parsing current URL:", err)
		return
	}
	if baseURL.Host != currentURL.Host {
		return
	}

	// Normalize the current URL
	normalizedURL, _ := NormalizeURL(currentURL.String())

	// Check if we've already crawled this page
	if count, exists := pages[normalizedURL]; exists {
		pages[normalizedURL] = count + 1
		return
	}

	// Add the current page to the map
	pages[normalizedURL] = 1

	// Print the current URL being crawled
	fmt.Printf("Crawling: %s\n", normalizedURL)

	// Get the HTML from the current URL
	htmlBody, err := GetHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
		return
	}

	// Get all URLs from the HTML
	urls, err := GetURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		return
	}

	// Recursively crawl each URL on the page
	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
