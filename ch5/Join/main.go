package main

import "fmt"

func Join(j string, str ...string) string {
    Str := str[0]
    for i, s := range str {
        if i != 0 {
            Str = Str + j + s
        }
    }
    return Str
}

func main() {
    s := Join("!!", "aa", "b", "c")
    fmt.Println(s)
}
