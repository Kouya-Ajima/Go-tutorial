package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Sqrt((x*x + y*y))
}

// 関数に関数を渡せる
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {

	// fmt.Println(hypot(5, 12)) // 13
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
