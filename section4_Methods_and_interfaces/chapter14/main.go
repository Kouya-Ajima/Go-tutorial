package main

import (
	"fmt"
)

func main() {

	// 空のインターフェースはなんの型でも格納できる。
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "Hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T) \n", i, i)
}
