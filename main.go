package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {

	argsWithProg := os.Args
	if len(os.Args) == 1 {
		fmt.Println("no website provided")
	}
	argsWithoutProg := os.Args[1:]

	fmt.Println("argsWithProg: ", argsWithProg)
	fmt.Println("argsWithoutProg", argsWithoutProg)
	fmt.Println("starting crawl of: ", os.Args[1])
	var rawURL = os.Args[1]
	var pages = make(map[string]int)
	baseURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing base URL:", err)
		return
	}
	var defaultMaxPage int = 25
	var defaultMaxConcurency int = 3
	if len(os.Args) > 2 {
		val, err := strconv.Atoi(os.Args[2])
		if err == nil {
			defaultMaxConcurency = val
		}
	}

	if len(os.Args) > 3 {
		val, err := strconv.Atoi(os.Args[3])
		if err == nil {
			defaultMaxPage = val
		}
	}
	maxPage := flag.Int("max", defaultMaxPage, "max pages to crawl")
	maxConcurency := flag.Int("c", defaultMaxConcurency, "max concurency")
	var wg sync.WaitGroup
	var mu sync.Mutex
	var concurencyControll = make(chan struct{}, *maxConcurency)
	var c = config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 &mu,
		concurrencyControl: concurencyControll,
		wg:                 &wg,
		maxPage:            *maxPage,
	}
	c.crawlPage(rawURL)
	wg.Wait()

	fmt.Println("crawl ended")
}
