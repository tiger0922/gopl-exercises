package main 

import "fmt"
import "gopl-exercises/ch2/tempconv"

func main() {
    fmt.Printf("Absolute Zero K = %v\n", tempconv.AbsoluteZeroK)
    fmt.Printf("Freezing Point K = %v\n", tempconv.FreezingK)
    fmt.Printf("Boiling Point K = %v\n", tempconv.CToK(tempconv.BoilingC))
    
}
