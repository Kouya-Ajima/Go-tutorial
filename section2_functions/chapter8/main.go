package main

import (
	"fmt"
	"math"
)

var eps float64 = math.Nextafter(1, 2) - 1

func Sqrt(x float64) (float64, int) {
	z := x / 2
	prev_z := z
	count := 0
	for i := 0; i < 10; i++ {
		count++
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev_z) <= eps {
			break
		}
		prev_z = z
	}
	return z, count
}

func main() {
	for i := 0; i < 10; i++ {
		x := float64(i + 1)
		z, count := Sqrt(x)
		w := math.Sqrt(x)
		fmt.Printf("loop count=%2v, Sqrt(%2v)=%18v, math.Sqrt(%2v)=%18v\n", count, x, z, x, w)
	}
}
