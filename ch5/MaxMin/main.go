// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import "fmt"

const (
    MaxInt = 1<<(32-1) - 1
    MinInt = -MaxInt - 1
)


//!+
func max(vals ...int) (int, error) {
    if len(vals) == 0 {
        return 0, fmt.Errorf("Please input at least one number.")
    }
	max := MinInt
	for _, val := range vals {
        if val > max {
            max = val
        }
	}
	return max, nil
}

func min(vals ...int) (int, error) {
    if len(vals) == 0 {
        return 0, fmt.Errorf("Please input at least one number.")
    }
	min := MaxInt
	for _, val := range vals {
        if val < min {
            min = val
        }
	}
	return min, nil
}
//!-


func main() {
	//!+main
    Print := func(num int, err error) {
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(num)
        }
    }
	Print(min())           //  "0"
	Print(max(3))          //  "3"
	Print(max(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	Print(min(values...)) // "10"
	//!-slice
}
