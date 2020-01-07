// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "Stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, filename := range counts {
        total := 0
        str, sep := "", "" 
        for name, n := range filename {
            total += n
            str += (sep + name + "[" + strconv.Itoa(n) + "]") 
            sep = " " 
        }
	    if total > 1 {
            fmt.Printf("%d\t%s\t(%s)\n", total, line, str)
	    }
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
        line := input.Text()
        if counts[line] == nil {
            counts[line] = make(map[string]int)
        }
        counts[line][filename]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
