package main

import "fmt"

func main() {
	// defer ステートメントは、deferへ私た関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させるもの
	// deferへ渡した関数の引数は、すぐに評価されますが、その関数自体は呼び出し元の関数がreturnするまで実行されない
	defer fmt.Println("world")

	fmt.Println("hello")
}
