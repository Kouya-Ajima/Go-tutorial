package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a, a[0], a[1])

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes, len(primes))

	// len(primes)で長さを取得
	var sum int
	for i := 0; i < len(primes); i++ {
		sum += primes[i]
	}
	fmt.Println(sum)
}
