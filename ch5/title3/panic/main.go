package main

import "fmt"

func weird() (ret string) {
    defer func() {
        ret = fmt.Sprintf("%v", recover()) // return omg
        ret = "checkmate"                  // return checkmate
    }()
    panic("omg")
}

func main() {
    fmt.Println(weird())
}
