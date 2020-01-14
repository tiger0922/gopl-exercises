// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import "fmt"
import "flag"

//!+
import "crypto/sha256"
import "crypto/sha512"

func main() {
    str1 := flag.String("s1", "x", "Input first string")
    str2 := flag.String("s2", "X", "Input second string")
    s1 := []byte(*str1)
    s2 := []byte(*str2)
	c1 := sha256.Sum256(s1)
	c2 := sha256.Sum256(s2)
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
    fmt.Printf("Difference: %d bits\n", checkDiff(c1, c2))
    sha512 := flag.Bool("sha512", false, "display sha512")
    sha384 := flag.Bool("sha384", false, "display sha384")
    flag.Parse()
    if *sha512 {
        mysha512(s1, s2)
    }
    if *sha384 {
        mysha384(s1, s2)
    }
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

func mysha512 (s1, s2 []byte) {
    c1 := sha512.Sum512(s1)
    c2 := sha512.Sum512(s2)
    fmt.Printf("\nsha512\nc1 = %x\nc2 = %x\n", c1, c2)
}

func mysha384 (s1, s2 []byte) {
    c1 := sha512.Sum384(s1)
    c2 := sha512.Sum384(s2)
    fmt.Printf("\nsha384\nc1 = %x\nc2 = %x\n", c1, c2)
}

func checkDiff(c1, c2 [32]uint8) int {
    count := 0
    for i := 0; i < 32; i++ {
        if byte(c1[i]) != byte(c2[i]) {
            count++
        }
    }
    return count    
}
//!-
