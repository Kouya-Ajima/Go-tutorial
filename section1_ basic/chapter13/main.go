package main

import (
	"fmt"
)

func main() {
	x, y := 3, 4
	f := float64(x)
	// f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}
