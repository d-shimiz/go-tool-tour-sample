package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// ifステートメントで宣言された変数は、elseブロック内でも使うことが可能
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3,2,10),
		pow(3,2,20),
	)
}
