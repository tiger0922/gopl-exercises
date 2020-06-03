package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(expand(os.Args[1], foo))
}

func expand(s string, f func(string) string) string {
    r := os.Args[2]
    return strings.Replace(s, r, f(r), len(s))
}

func foo(s string) string {
    str := "^" + s + "^"
    return str
}
