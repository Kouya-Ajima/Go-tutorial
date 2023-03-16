package main

import (
	"fmt"
)

// x, yという名前を定義しながらの関数
// naked returnステートメントは、短い関数でのみ利用すべき
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// x, y が返る
	fmt.Println(split(17))
}
