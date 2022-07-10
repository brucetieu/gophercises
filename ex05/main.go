package main

import (
	"flag"
	"fmt"

	"github.com/brucetieu/gophercises/ex05/sitemap"
)


func main() {
	urlPtr := flag.String("url", "", "The url to build a sitemap out of")

	flag.Parse()

	sitemap.URL = *urlPtr

	fmt.Println(sitemap.BuildSitemap(*urlPtr))

}