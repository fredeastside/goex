package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println(CountWordsAndImages("https://yandex.ru"))
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(doc *html.Node) (words, images int) {

	if doc.Type == html.ElementNode && doc.Data == "img" {
		images++
	}
	if doc.Type == html.TextNode {
		words += len(strings.Split(doc.Data, " "))
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		i, w := countWordsAndImages(c)
		images += i
		words += w
	}

	return
}
