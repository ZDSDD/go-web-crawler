package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPage            int
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.mu.Lock()
	if cfg.maxPage <= len(cfg.pages) {
		fmt.Println("max page limit exceeded.")
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	// Check if the current URL is on the same domain as the base URL
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Error parsing current URL:", err)
		return
	}
	if cfg.baseURL.Host != currentURL.Host {
		return
	}

	// Normalize the current URL
	normalizedURL, err := normalizeURL(currentURL.String())
	if err != nil {
		fmt.Println("Couldn't normalize: ", currentURL.String())
		return
	}
	if !cfg.addPageVisit(normalizedURL) {
		return
	}
	// Print the current URL being crawled
	fmt.Printf("Crawling: %s\n", normalizedURL)

	// Get the HTML from the current URL
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
		return
	}

	// Get all URLs from the HTML
	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		return
	}

	// Recursively crawl each URL on the page
	for _, url := range urls {
		cfg.wg.Add(1)
		go func() {
			defer cfg.wg.Done()
			cfg.concurrencyControl <- struct{}{}
			cfg.crawlPage(url)
			<-cfg.concurrencyControl
		}()
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {

	// Check if we've already crawled this page
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if count, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL] = count + 1
		return false
	}

	// Add the current page to the map
	cfg.pages[normalizedURL] = 1
	return true
}
