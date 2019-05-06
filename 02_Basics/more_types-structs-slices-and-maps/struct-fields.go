package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// 構造体を変数定義
	v := Vertex{1,2}
	// 構造体のフィールドへアクセス
	v.X = 4
	fmt.Println(v.X)
}
