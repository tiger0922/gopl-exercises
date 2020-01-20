package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	freq := make(map[string]int)
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		os.Exit(1)
	}

	defer file.Close()

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		parts := strings.Trim(word, ",.\\[]-():;`'\"/")
		freq[parts]++
	}
	fmt.Printf("word\tcount\n")
	for c, n := range freq {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
