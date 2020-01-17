package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		fmt.Println(s)
		s = string(reverse([]byte(s)))
		fmt.Println(s)
		fmt.Println(string(squash([]byte(s))))
	}
}

func squash(bytes []byte) []byte {
	var out []byte
	var last rune

	// Reference : https://github.com/xingdl2007/gopl-solutions/blob/master/ch4/4.6/main.go

	for i := 0; i < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])

		if !unicode.IsSpace(r) {
			out = append(out, bytes[i:i+size]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
			out = append(out, ' ')
		}

		last = r
		i += size
	}

	return out
}

func reverse(bytes []byte) []byte {
	// Change runes which are greater than 1bit first
	for i := 0; i < len(bytes); {
		_, size := utf8.DecodeRune(bytes[i:])
		s := size

		if size > 1 {
			for j := 0; s > j+1 && j < 2; j, s = j+1, s-1 {
				bytes[i+j], bytes[i+s-1] = bytes[i+s-1], bytes[i+j]
			}
		}
		i += size
	}
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}
