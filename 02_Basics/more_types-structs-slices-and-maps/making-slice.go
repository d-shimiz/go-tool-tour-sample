package main

import "fmt"

func main() {
	// スライスは、組み込みのmake関数を使用して作成することができる。
	// これは動的サイズの配列を作成する方法
	// make関数はゼロ化された配列を割り当て、その配列を指すスライスを返す。
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
