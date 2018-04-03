package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse error: %v\n", err)
	}

	m := make(map[string]int)
	visit(m, doc)

	for el, count := range m {
		fmt.Println(el, " - ", count)
	}
}

func visit(m map[string]int, doc *html.Node) {

	if doc == nil {
		return
	}

	if doc.Type == html.TextNode && doc.Parent.Data != "script" && doc.Parent.Data != "style" {
		m[doc.Data]++
	}

	visit(m, doc.FirstChild)
	visit(m, doc.NextSibling)
}
