package main

import (
    "fmt"
    "math"
)

func main() {
    // Goでは、最初の文字が大文字で始まる名前は、外部のパッケージから参照できるエクスポートされた名前である。
    // Piはmathパッケージでエクスポートされている。
    // fmt.Println(math.pi)
    fmt.Println(math.Pi)
}

