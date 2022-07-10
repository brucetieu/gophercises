package main

import (
	"flag"
	"fmt"

	"github.com/brucetieu/gophercises/ex05/sitemap"
)


func main() {
	urlPtr := flag.String("url", "", "The url to build a sitemap out of")
	maxDepth := flag.Int("maxdepth", 5, "Define the maximum number of links to follow")

	flag.Parse()

	sitemap.URL = *urlPtr
	sitemap.MaxDepth = *maxDepth
	
	fmt.Println(sitemap.BuildSitemap())

}