package main

import "fmt"

func main() {
	// 配列は固定長、スライスは可変長
	// より柔軟な配列とも言える。
	// 型 []T は Tのスライスを表す。
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// コロンで区切られた2つのインデックス low と high の境界を指定することでスライスが形成される。
	var s []int = primes[1:4]
	fmt.Println(s)
}
