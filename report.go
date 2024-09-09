package main

import (
	"fmt"
)

func printRaport(pages map[string]int, baseUrl string) {
	fmt.Print("\nREPORT for ", baseUrl, "\n")
	fmt.Print("URL: ", baseUrl, "\n")
	fmt.Print("$$$$$$$$$$$$$$$$$$$$\n")
	for k, v := range pages {
		fmt.Printf("Found %d internal links to %s\n", v, k)
	}
}
