package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i<10; i++ {
		// deferへ渡した関数が複数ある場合、その呼び出しはスタック(stack)される。
		// 呼び出し元の関数がreturnする時、deferへ渡した関数はLIFOの順番で実行される。
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

