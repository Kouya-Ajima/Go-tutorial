package main

import "fmt"

// 関数の引数リストと同じ原理で型を最後に持ってきて省略できる
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
