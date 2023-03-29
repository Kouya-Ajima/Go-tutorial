package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64

// ////// インターフェース実装 //////////
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//////////////////////////////////////

func main() {

	var a, b Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	//　型を型にいれる。
	a = f
	b = &v
	// Error -> v *Vertex でしか Abs()が実装されていないため。
	// a = v

	// インターフェースに実装されている関数を使用しなければエラーになる。
	fmt.Println(a.Abs(), "\n", b.Abs())
}
