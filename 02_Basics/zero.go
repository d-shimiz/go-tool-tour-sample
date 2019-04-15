package main

import "fmt"

func main() {
    // 変数に初期値を与えず宣言すると、ゼロ値が与えられる。
    // ゼロ値は型によって以下のように与えられる。
    //   整数型(int, floatなど): 0
    //   bool型: false
    //   string型: "" (空文字列(empty string))
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v, %v, %v, %q\n", i, f, b ,s)
}
