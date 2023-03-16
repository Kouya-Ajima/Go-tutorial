package main

import "fmt"

// 返る引数の方をかく。
func swap(x, y string) (string, string) {
	return y, x
}

func main() {

	a, b := swap("222", "111")
	fmt.Println(a, b)
}
