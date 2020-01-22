package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "summary: %v\n", err)
		os.Exit(1)
	}
	var mapping = make(map[string]int)
	summary(mapping, doc)

	for k, v := range mapping {
		fmt.Printf("%-10s\t%d\n", k, v)
	}

}

func summary(mapping map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		mapping[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		summary(mapping, c)
	}

}
