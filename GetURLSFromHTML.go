package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var urls []string
	var traverseNodes func(*html.Node)
	traverseNodes = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					parsedURL, err := url.Parse(attr.Val)
					if err != nil {
						continue
					}
					resolvedURL := baseURL.ResolveReference(parsedURL)
					urls = append(urls, resolvedURL.String())
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverseNodes(c)
		}
	}

	traverseNodes(doc)
	return urls, nil
}
