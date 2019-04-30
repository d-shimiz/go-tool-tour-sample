package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// ifステートメントは、forのように、条件の前に、評価するための簡単なステートメントを書くことができる
	// ここで宣言された変数は、ifのスコープ内だけで有効
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}
