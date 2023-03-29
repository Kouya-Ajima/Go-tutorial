package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

var eps float64 = math.Nextafter(1, 2) - 1

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := 1.0
	prev_z := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev_z) <= eps {
			break
		}
		prev_z = z
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
