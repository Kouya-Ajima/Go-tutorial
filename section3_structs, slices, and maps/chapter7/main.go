package main

import (
	"fmt"
)

func main() {
	primes := [6]int{1, 2, 3, 4, 5, 6}
	// １からは表示される。 4番目 は表示されない。
	var s []int = primes[1:4]
	var s1 []int = primes[:4]
	var s2 []int = primes[1:]

	fmt.Println(s, s1, s2)
}
