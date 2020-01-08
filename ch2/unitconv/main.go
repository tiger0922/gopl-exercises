package main

import (
	"fmt"
	"os"
	"strconv"
    "bufio"

	"gopl-exercises/ch2/tempconv"
	"gopl-exercises/ch2/weightconv"
	"gopl-exercises/ch2/lengthconv"
)

func main() {
    if len(os.Args) < 2 {
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
		    t, err := strconv.ParseFloat(scanner.Text(), 64)
            if err != nil {
                fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
                os.Exit(1)
            }
            conv(t) 
        } 
    }
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
			os.Exit(1)
		}
        conv(t)
	}
}

func conv(t float64) {
    f := tempconv.Fahrenheit(t)
    c := tempconv.Celsius(t)
    fmt.Printf("%s = %s, %s = %s\n",
        f, tempconv.FToC(f), c, tempconv.CToF(c))

    ft := lengthconv.Feet(t)
    m  := lengthconv.Meter(t)
    fmt.Printf("%s = %s, %s = %s\n",
        ft, lengthconv.FtToM(ft), m, lengthconv.MToFt(m))

    lb := weightconv.Pound(t)
    kg := weightconv.Kilogram(t)
    fmt.Printf("%s = %s, %s = %s\n",
        lb, weightconv.LbToKg(lb), kg, weightconv.KgToLb(kg))

}
