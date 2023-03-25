package main

import (
	"fmt"
)

func printSlice(s string, x []int) {
	fmt.Printf(
		"%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {

	// Capのみ
	a := make([]int, 5)
	printSlice("a", a)

	// Len、Cap
	b := make([]int, 0, 5)
	printSlice("b", b)

	// これはできない。
	// X := make([]int)
	// printSlice("X", X)

	c := b[:2]
	printSlice("c", c)

	// ｃからスライスを生成できなくても、
	// キャパシティの容量内なら生成可能
	d := c[2:5]
	printSlice("d", d)

	// これはCap以上にスライスを取得しようとしているからエラー
	x := c[2:9]
	printSlice("x", x)

}
