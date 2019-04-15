package main

import "fmt"

// 関数は複数の戻り値を返すことができる。
// 以下のswap関数は2つのstringを返す。
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a,b)
}

