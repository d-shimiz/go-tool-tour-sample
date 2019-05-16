package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// for ループを利用する range は、スライスや、マップ(map)を一つずつ反復処理するために使う
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
