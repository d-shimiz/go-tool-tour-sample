package main

import "fmt"

func main() {
	sum := 1
	// セミコロン(;)を省略することも可能。C言語などにあるwhileはGoではforだけを使う
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
