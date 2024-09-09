package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// get body
	resp, err := http.Get(rawURL)
	if err != nil {
		fmt.Printf("getHTML: error when fetching: %s\n", rawURL)
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", fmt.Errorf("couldn't fetch %s. Status code: %d", rawURL, resp.StatusCode)
	}
	if contentType := resp.Header.Get("Content-Type"); contentType == "" {
		return "", fmt.Errorf("there is no Content-Type in the response header %s", rawURL)
	} else if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("content type is not text/html. Got %s instead", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(body), nil
}
