package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {

	s := []int{2, 4, 5, 6, 41}
	printSlice(s)

	//	スライスが取得されない（0）のときは値が更新されない。
	// この場合、S にスライスの値は格納されない。
	s = s[:0]
	printSlice(s)

	// この時点で初めて S に値が格納
	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}
