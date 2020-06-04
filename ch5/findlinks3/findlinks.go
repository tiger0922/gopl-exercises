// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"
    "io"
    "net/http"
    "path"
    "strings"

	"gopl-exercises/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
    fmt.Println(local)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
    s := strings.SplitAfter(url, "/")
    domain := strings.Join(s[:3], "")
    if domains[domain] {
        fetch(url)
    }
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl
var domains map[string]bool
//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
    domains = make(map[string]bool)
    for i := 1; i < len(os.Args); i++ {
        domains[os.Args[i]] = true
    }
	breadthFirst(crawl, os.Args[1:])
}

//!-main
