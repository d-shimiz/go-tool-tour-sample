package main

import "fmt"

func main() {
    // 明示的な型を指定せずに変数を宣言する場合( := や var = のいずれか)、変数の型は右側の変数から型推論される。
    // 右側に型を指定しない数値である場合、左側の新しい変数は右側の定数の精度に基いて int, float64, complex128 の型になる。
    v := 42 // change me!
    // v := 3.14 // change me!
    // v := 0.5 + 0.5i  // change me!
    fmt.Printf("v is of type %T\n", v)
}

