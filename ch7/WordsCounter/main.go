// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//!+wordcounter

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
    scanner := bufio.NewScanner(bytes.NewReader(p))
    scanner.Split(bufio.ScanWords)
    l := 0
    for scanner.Scan() {
        l++
    }
    *c += WordCounter(l)
    return l, nil
}
func (c *LineCounter) Write(p []byte) (int, error) {
    scanner := bufio.NewScanner(bytes.NewReader(p))
    scanner.Split(bufio.ScanLines)
    l := 0
    for scanner.Scan() {
        l++
    }
    *c += LineCounter(l)
    return l, nil
}

//!-wordcounter

func main() {
    //!+main
    var c WordCounter
    c.Write([]byte("hello 1 2 3 4"))
    fmt.Println(c) // "5", = len("hello")

    c = 0 // reset the counter
    var name = "Dolly"
    fmt.Fprintf(&c, "hello, %s", name)
    fmt.Println(c) // "12", = len("hello, Dolly")
    //!-main

    var l LineCounter
    l.Write([]byte("Yo\n1\n2"))
    fmt.Println(l)
}
