package sitemap

import (
	"container/list"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/brucetieu/gophercises/ex04/parser"
)

var (
	URL string
	MaxDepth int
	HTTPPrefix = "https://"
	SlashPrefix = "/"

	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

type SitemapURL struct {
	Loc string `xml:"loc"`
}

type SitemapURLset struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNs   string   `xml:"xmlns,attr"`
	XMLLoc  []SitemapURL `xml:"url"`
}

func GetHTML(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(html)
}

func BuildSitemap() string {
	
	sitemapUrlset := &SitemapURLset{}
	sitemapUrlset.XMLNs = "http://www.sitemaps.org/schemas/sitemap/0.9"
	sitemapUrls := make([]SitemapURL, 0)

	visited := make(map[string]bool, 0)

	siteUrls := TraversePage(URL, visited)
	for _, siteUrl := range siteUrls {
		sitemapUrls = append(sitemapUrls, SitemapURL{siteUrl})
	}

	sitemapUrlset.XMLLoc = sitemapUrls

	out, err := xml.MarshalIndent(sitemapUrlset, "", "\t")
	if err != nil {
		panic(err)
	}

	final := []byte(xml.Header + string(out))
	return string(final)
}

func TraversePage(url string, visited map[string]bool) []string {
	queue := list.New()
	pages := make([]string, 0)

	queue.PushBack(url)

	for queue.Len() > 0 {
		
		// Return MaxDepth number of links
		if len(pages) == MaxDepth {
			return pages
		}

		e := queue.Front()
		linkEle := queue.Remove(e)

		link := linkEle.(string)
		
		if _, ok := visited[link]; !ok {
			visited[link] = true
			pages = append(pages, link)

			htmlPage := GetHTML(link)
			newLinks := FilteredLinks(htmlPage)

			for _, newLink := range newLinks {
				queue.PushBack(newLink)
			}
		}
	}
	
	return pages
}

func FilteredLinks(htmlPage string) []string {
	links := parser.HTMLParser(htmlPage)
	filteredLinks := make([]string, 0)

	for _, link := range links {
		if strings.HasPrefix(link.Href, HTTPPrefix) {
			if checkDomain(link.Href) {
				filteredLinks = append(filteredLinks, link.Href)
			}
		} else if strings.HasPrefix(link.Href, SlashPrefix) {
			filteredLinks = append(filteredLinks, URL + link.Href)
		} 
		
	}

	return filteredLinks

}

func checkDomain(url string) bool {
	return strings.HasPrefix(url, URL)
}
