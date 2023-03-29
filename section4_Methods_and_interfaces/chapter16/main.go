package main

import (
	"fmt"
)

// 何でもインターフェースを引数に取る
func do(i interface{}) {

	// type で型名を取得する
	// 型に応じたSwitch文
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("Hello")
	do(true)
}
