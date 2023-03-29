package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs(), AbsFunc(v))

	p := &Vertex{3, 4}
	// p を引数として渡すときは、*p のようにポインタを渡さないとエラーになる
	// p.Abs() は (*p).Abs() として解釈される。
	fmt.Println(p.Abs(), AbsFunc(*p))
}
