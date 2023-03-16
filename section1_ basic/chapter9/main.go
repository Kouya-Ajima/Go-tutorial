package main

import (
	"fmt"
)

var i, j int = 1, 2

func main() {
	// 使用していないとVarはエラーになる。
	// Var は　(動的型付けになる)
	var c, java, python = true, false, "No!"
	fmt.Println(i, j, c, java, python)
}
