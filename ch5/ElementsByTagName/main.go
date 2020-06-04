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
    tags := os.Args[2:]
    resp, err := http.Get(url)

    if err != nil {
        return
    }

    defer resp.Body.Close()
    doc, err := html.Parse(resp.Body)

    if err != nil {
        return
    }

    for _, n := range ElementsByTag(doc, tags...) {
        fmt.Printf("+%v\n", n)
    }

    return
}

func ElementsByTag(n * html.Node, tags ...string) []*html.Node {
    nodes := make([]*html.Node, 0)
    keep := make(map[string]bool, len(tags))
    for _, t := range tags {
        keep[t] = true
    }

    pre := func(n *html.Node) bool {
        if n.Type != html.ElementNode {
            return true
        }
        _, ok := keep[n.Data]
        if ok {
            nodes = append(nodes, n)
        }
        return true
    }
    forEachElement(n, pre, nil)
    return nodes
}

func forEachElement(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
    u := make([]*html.Node, 0) // unvisited
    u = append(u, n)
    for len(u) > 0 {
        n = u[0]
        u = u[1:]
        if pre != nil {
            if !pre(n) {
                return n
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            u = append(u, c)
        }
        if post != nil {
            if !post(n) {
                return n
            }
        }
    }
    return nil
}

