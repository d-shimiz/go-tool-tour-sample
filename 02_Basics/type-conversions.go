package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y int = 3, 4
    // 変数 v 、型 T があった場合、 T(v) は、変数 v を T 型へ変換する。
    var f float64 = math.Sqrt(float64(x*y + y*y))
    var z uint = uint(f)
    fmt.Println(x, y ,z)
}
