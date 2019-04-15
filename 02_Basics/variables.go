package main

import "fmt"

// varステートメントは変数(variables)を宣言する。
// 関数の引数リストと同様に、複数の変数の最後に型を書くことで、変数のリストを宣言できる。
var c, python, java bool

func main() {
    var i int
    fmt.Println(i, c, python, java)
}
