package main

import "fmt"

func main() {
	m := make(map[string]int)

	// mへの要素の挿入や更新 m[key]=elem
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// meへの要素の削除 delete(m, key)
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// キーに対する要素が存在するかどうかは、elem, ok = m[key] で確認できる
	// m に key があれば変数はokとなり、存在しなければ、okはfalseとなる。
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

