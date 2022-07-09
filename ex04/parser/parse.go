package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ReadHTML(filename string) string {
	html, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(html)
}

func PrettyFormat(data interface{}) string {
	bytes, _ := json.MarshalIndent(data, "", "\t")
	return string(bytes)
}

func traverseHTML(node *html.Node, links *[]Link) {
	if node.Type == html.ElementNode && node.Data == "a" {
		href := ""
		for _, a := range node.Attr {
			if a.Key == "href" {
				href = a.Val
				break
			}
		}

		atagText := traverseTag(node, "")
		if href != "" {
			link := Link{
				Href: href,
				Text: atagText,
			}
			*links = append(*links, link)
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverseHTML(c, links)
	}
}

func HTMLParser(htmlString string) []Link{
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		log.Fatal(err)
	}

	links := make([]Link, 0)

	traverseHTML(doc, &links)
	
	return links
}

func traverseTag(node *html.Node, s string) string {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			s += strings.TrimSpace(c.Data) + " "
			
		}
		s = traverseTag(c, s)
	}

	return strings.TrimSpace(s)
}