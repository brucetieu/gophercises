package main

import (
	"flag"
	"fmt"

	"github.com/brucetieu/gophercises/ex04/parser"
)


func main() {
	filePtr := flag.String("htmlfile", "", "HTML file to parse")

	flag.Parse()

	htmlString := parser.ReadHTML(*filePtr)
	links := parser.HTMLParser(htmlString)

	fmt.Println("links: ", parser.PrettyFormat(links))
}
