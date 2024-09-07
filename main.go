package main

import (
	"fmt"
	"log"
	"os"
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
	htmlBody, err := GetHTML(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("fetch ended successfuly!")
	fmt.Println(htmlBody)
}
