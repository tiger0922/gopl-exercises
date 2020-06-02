package main

import (
    "fmt"
    "os"
    "strings"
    "net/http"

    "golang.org/x/net/html"
)

func main() {
    words, images, _ := CountWordsAndImages(os.Args[1])
    fmt.Printf("Words:  %d\nImages:  %d\n", words, images)
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
    }
    words, images = countWordsAndImages(doc)
    return words, images, err
}

func countWordsAndImages(n *html.Node) (words, images int) {
    if n.Type == html.TextNode {
        w := strings.Trim(n.Data, ",.\\[]-():;`'\"/")
        word := len(w)
        words += word
    } else if n.Data == "img" {
        images++
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        word, image := countWordsAndImages(c)
        words += word
        images += image
    }
    return words, images
}
