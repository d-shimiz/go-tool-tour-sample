package main

import (
	"fmt"
	"math"
)

// Newton's method
func Sqrt(x float64) float64 {
	// x := 2.0
	z := 1.0
	n := 0.0
	for i := 1.0 ; i < 10; i++ {
		z -= ( z * z - x) / ( 2 * z )
		if math.Abs( z - n ) < 0.0000001 {
			fmt.Println(i)
			return z
		}
		n = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
