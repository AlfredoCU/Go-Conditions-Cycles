package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaration of variable
	var e float64
	//var f float64 = 1

	// Calculation of Number e with factorial
	//for i := 1; i <= 10; i++ {
	//	f *= float64(i)
	//	e = 1 / f
	//}

	// Calculation of Number e
	for n := 1; n <= 1000000; n++ {
			e = math.Pow(1 + (1 / float64(n)), float64(n))
	}

	// Print result
	fmt.Print(e) // ê = 2.7182804690957534
				 // e = 2.718281828459045
}