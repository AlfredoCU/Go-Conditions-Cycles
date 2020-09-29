package main

import (
	"fmt"
	// "math"
)

func main() {
	// Declaration of variable
	var e float64
	var f float64 = 1

	// Calculation of Number e with factorial
	for i := 0; i < 18; i++ { // 17 Iterations for ê
		f = 1
		for j := i; j > 1; j-- {
			f *= float64(j)
		}
		e += 1 / f
	}

	// Calculation of Number e
	// for n := 0; n <= 1000000; n++ {
	//		e = math.Pow(1 + (1 / float64(n)), float64(n))
	// }

	// Print result
	fmt.Print(e)
	// ê = 2.7182818284590455 // One method
	// ê = 2.7182804690957534 // Two method
	// e = 2.718281828459045
}