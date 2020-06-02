package main

import (
    "fmt"
    "os"
    "strings"

    "golang.org/x/net/html"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "summary: %v\n", err)
        os.Exit(1)
    }
    printText(doc)
}

func printText(n *html.Node) {
    if n.Data == "script" || n.Data == "style" {
        return
    }
    if n.Type == html.TextNode {
        text := n.Data
        text = strings.Trim(text, "\n ")
        if(text != "") {
            fmt.Println(text)
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        printText(c)
    }
}

