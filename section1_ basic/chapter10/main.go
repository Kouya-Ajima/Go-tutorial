package main

import (
	"fmt"
)

func main() {
	// 関数内なら、Varを省略してかける。(動的型付けになる)
	k := 3
	f := "2"

	var i, j int = 1, 2
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
	fmt.Printf("Type: %T\n Type: %T", k, f)
}
