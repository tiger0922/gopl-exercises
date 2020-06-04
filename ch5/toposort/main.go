// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
    "fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
    "algorithms":       {"data structures"},
    "calculus":         {"linear algebra"},
    "linear algebra":   {"calculus"},

    "compilers": {
        "data structures",
        "formal languages",
        "computer organization",
    },

    "data structures":       {"discrete math"},
    "databases":             {"data structures"},
    "discrete math":         {"intro to programming"},
    "formal languages":      {"discrete math"},
    "networks":              {"operating systems"},
    "operating systems":     {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
    depth = 0
    for i, course := range topoSort(prereqs) {
        fmt.Printf("%d:\t%s\n", i+1, course)
    }
}

var depth int

func topoSort(m map[string][]string) []string {
    var order []string
    seen := make(map[string]bool)
    topdown := make(map[string]bool)
    var visitAll func(items []string, topdown map[string]bool)

    visitAll = func(items []string, topdown map[string]bool) {
        for _, item := range items {
            if depth == 0 {
                topdown = make(map[string]bool)
            }
            if !seen[item] {
                seen[item] = true
                topdown[item] = true
                depth++
                visitAll(m[item], topdown)
                depth--
                order = append(order, item)
            } else if topdown[item] {
                fmt.Printf("Cycle detected at %s\n", item)
            }
        }
    }

    for k := range m {
        visitAll([]string{k}, topdown)
    }

    return order
}

//!-main
