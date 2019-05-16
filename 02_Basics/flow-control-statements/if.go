package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// ifステートメントでは、forループと同様に初期化、条件式、後処理ステートメントの括弧()は不要
	// 中括弧{}は必要
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
