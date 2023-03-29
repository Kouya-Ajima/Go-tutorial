package main

import (
	"fmt"
)

func main() {

	var i interface{}
	i = "Hello"

	// アサーションで型を定められる
	s := i.(string)
	fmt.Println(s)

	// 型があっているかどうかの判断
	s, ok := i.(string)
	fmt.Println(s, ok)

	// すでに入っている ｉ の変数 と、違う型を定めた場合
	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f) // error: panic

}
