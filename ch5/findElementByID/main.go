// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
    "fmt"
    "net/http"
    "os"

    "golang.org/x/net/html"
)


func main(){
    url := os.Args[1]
    id := os.Args[2]
    resp, err := http.Get(url)

    if err != nil {
        return
    }

    defer resp.Body.Close()
    doc, err := html.Parse(resp.Body)

    if err != nil {
        return
    }

    fmt.Printf("%+v\n", ElementByID(doc, id))

    return
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func ElementByID (doc *html.Node, id string) *html.Node {
    return forEachNode(id, doc, startElement, endElement)
}
func forEachNode(id string, n *html.Node, pre, post func(n *html.Node, id string) bool) *html.Node {
    if pre != nil {
        if !pre(n, id) {
            return n
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        node := forEachNode(id, c, pre, post)
        if forEachNode(id, c, pre, post) != nil {
            return node
        }
    }

    if post != nil {
        if !post(n, id) {
            return n
        }
    }
    return nil
}

//!-forEachNode

//!+startend

func startElement(n *html.Node, id string) bool {
    if n.Type == html.ElementNode {
        for _, a := range n.Attr {
            if a.Key == "id" && a.Val == id {
                return false
            }
        }
    }
    return true
}

func endElement(n *html.Node, id string) bool {
    if n.Type == html.ElementNode {
        for _, a := range n.Attr {
            if a.Key == "id" && a.Val == id {
                return false
            }
        }
    }
    return true
}

//!-startend
