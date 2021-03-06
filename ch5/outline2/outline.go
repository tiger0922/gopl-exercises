// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"

    "golang.org/x/net/html"
)

func main() {
    for _, url := range os.Args[1:] {
        outline(url)
    }
}

func outline(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)
    if err != nil {
        return err
    }

    var depth int

    startElement := func(n *html.Node) {
        if n.Type == html.ElementNode {
            str := fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
            for _, v := range n.Attr {
                str += fmt.Sprintf(` %s="%s"`, v.Key, v.Val)
            }
            if n.FirstChild == nil {
                str += "/"
            }
            str += ">"
            fmt.Println(str)
            depth++
        } else if n.Type == html.TextNode {
            text := strings.TrimSpace(n.Data)
            texts := strings.Split(text, "\n")
            space := fmt.Sprintf("\n%*s", depth*2, "")
            text = strings.Join(texts, space)
            if len(text) != 0 {
                fmt.Printf("%*s%s\n", depth*2, "", text)
            }
        } else if n.Type == html.CommentNode {
            fmt.Printf("<!--%s-->\n", n.Data)
        }
    }

    endElement := func(n *html.Node) {
        if n.Type == html.ElementNode {
            depth--
            if n.FirstChild != nil {
                fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
            }
        }
    }

    //!+call
    forEachNode(doc, startElement, endElement)
    //!-call

    return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    if pre != nil {
        pre(n)
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }

    if post != nil {
        post(n)
    }
}

//!-forEachNode

//!+startend

//!-startend
