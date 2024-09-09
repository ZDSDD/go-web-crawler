package main

import (
	"fmt"
	"net/url"
	"os"
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
	if len(argsWithoutProg) > 1 {
		fmt.Printf("too many arguments provided")
		os.Exit(1)
	}
	fmt.Println("starting crawl of: ", os.Args[1])
	var rawURL = os.Args[1]
	var pages = make(map[string]int)
	baseURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing base URL:", err)
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var concurencyControll = make(chan struct{}, 10)
	var c = config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 &mu,
		concurrencyControl: concurencyControll,
		wg:                 &wg,
	}
	c.crawlPage(rawURL)
	wg.Wait()

	fmt.Println("crawl ended")
}
