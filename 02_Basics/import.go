package main

// 括弧でパッケージのインポートをグループ化し、factoredインポートステートメント(factored import statement)としている。
import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("Now you have %g problem.", math.Sqrt(7))
}

