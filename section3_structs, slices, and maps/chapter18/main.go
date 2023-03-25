package main

import "fmt"

func Pic(dy, dx int) [][]uint8 {
	s := make([][]uint8, dy)
	fmt.Println("s:", s)

	for i := range s {
		s[i] = make([]uint8, dx)
		fmt.Printf("s[%v]: %v\n", i, s[i])
	}

	fmt.Println("s:", s)

	return s
}

func main() {
	Pic(4, 7)
}
