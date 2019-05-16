package main

import (
    "fmt"
    "math/cmplx"
)

// Goの基本型は以下の通り。
// bool
// 
// string
// 
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// 
// byte // uint8 の別名
// 
// rune // int32 の別名
//      // Unicode のコードポイントを表す
// 
// float32 float64
//
// complex64 complex128

var (
    ToBe    bool       = false
    MaxInt  uint64     = 1 << 64 - 1
    z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    fmt.Printf("Type: %T value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T value: %v\n", z, z)
}

// int, uint, uintptr 型は、32-bitのシステムでは32 bitで、64-bitのシステムでは64 bitとなる。
// サイズ、符号なし( unsigned )整数の型を使うための特別な理由がない限り、整数の変数が必要な場合は int を使うようにしよう。
