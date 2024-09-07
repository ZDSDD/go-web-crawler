package main

import (
	"strings"
)

func normalizeURL(url string) (string, error) {
	// Remove protocol (http:// or https://)
	url = strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "http://")

	// Remove www. if present
	url = strings.TrimPrefix(url, "www.")

	// Remove trailing slash if present
	url = strings.TrimSuffix(url, "/")

	return url, nil
}
