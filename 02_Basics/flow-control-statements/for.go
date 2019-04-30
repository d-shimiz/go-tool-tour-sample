package main

import "fmt"

func main() {
	sum := 0
	// 基本的にforループはセミコロン; で3つの部分に分かれている
	// 初期化ステートメント: 最初のイテレーション(繰り返し)の前に初期化が実行される
	// 条件式: イテレーション毎に評価される
	// 後処理ステートメント: イテレーション毎の最後に実行される
	// 初期化ステートメントは、短い変数宣言によく利用される。その変数はforステートメントのスコープ内でのみ有効1
	// ループは、条件式の評価がfalseとなった場合にイテレーションを停止する。
	for i := 0; i<10; i++ {
		sum += i
	}
	fmt.Println(sum)
}