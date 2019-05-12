package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// img は要素数==容量のuint8型の配列への参照
	// make(型, 要素数, 容量)として使う。
	// スライスが初期化され、デフォルトではゼロ値として0が格納されている。
	// 容量を省略すると、要素数==容量となる。
	img_data := make([][]uint8, dy)

	// 配列に対するrangeでは、index番号と要素の値が返される。
	for y := range img_data {
		img_data[y] = make([]uint8, dx)
		for x := range img_data[y] {
			img_data[y][x] = uint8(x * y)
		}
	}

	return img_data
}

func main() {
	pic.Show(Pic)
}

