package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	// runtime.GOOSで実行中のOSの情報を取得する。変数osの内容によって分岐
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, opnebsd
		// @lan9, windows...
		fmt.Printf("%s.", os)
	}
}
