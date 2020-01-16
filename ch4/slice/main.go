package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
)

func main () {
	input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        s := strings.Fields(input.Text())
	    fmt.Printf("%s\n", s)
        i := 0
        for {
            if (s[i] == s[i+1]) {
                s = append(s[:i], s[i+1:]...)
            } else {
                i++
            }
            if i == len(s) - 1 {
                break
            }
        }
	    fmt.Printf("%s\n", s)
    }
}
