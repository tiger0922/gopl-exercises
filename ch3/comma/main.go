// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
    "fmt"
    "os"
    "sort"
)

func main() {
    for i := 1; i < len(os.Args); i++ {
        fmt.Printf("  %s\n", comma(os.Args[i]))
    }
    fmt.Printf("  %t\n", anagram(os.Args[1], os.Args[2]))
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
    n := len(s)
    for i := n-1; i > 0; i-- {
        if s[i] == '.' {
            n = i
            break
        }
    }
    /* Recursive Method
    if n <= 3 {
        return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]
    */
    for i := n; i > 0; i-=3 {
        if i-3 > 0 {
            s = s[:i-3] + "," + s[i-3:]
        }
    }
    return s
}

//My method : sorting and compare
func anagram(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    n1 := make([]int, len(s1))
    n2 := make([]int, len(s2))

    for i := 0; i < len(s1); i++ {
        n1[i] = int(s1[i])-int('0')
        n2[i] = int(s2[i])-int('0')
    }
    sort.Ints(n1)
    sort.Ints(n2)
    for i := 0; i < len(s1); i++ {
        if n1[i] != n2[i] {
            return false
        }
    }
    return true

}

func isAnagram(a, b string) bool {
    aFreq := make(map[rune]int)
    for _, c := range a {
        aFreq[c]++
    }
    bFreq := make(map[rune]int)
    for _, c := range b {
        bFreq[c]++
    }
    for k, v := range aFreq {
        if bFreq[k] != v {
            return false
        }
    }
    for k, v := range bFreq {
        if aFreq[k] != v {
            return false
        }
    }
    return true
}
//!-
