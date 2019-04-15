package main

import "fmt"

// 関数は0個以上の引数を取ることができる。
// この例では、add関数は、int型の2つのパラメータを取る。
// 変数名の後ろに型名を書くことに注意する。
func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42,13)})
}

