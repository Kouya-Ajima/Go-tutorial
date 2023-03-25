package main

import (
	"fmt"
)

func main() {

	names := [4]string{
		"john", "paul", "georege", "ringo",
	}
	fmt.Println(names)

	a := names[:2]
	b := names[1:3]
	fmt.Println(a, b)

	// 元の配列に変更が反映される
	b[0] = "XXXXX"
	fmt.Println(a, b)
	fmt.Println(names)

}
