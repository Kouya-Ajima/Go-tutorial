package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// %+v で、変数名と値を表示する
	v := Vertex{3, 4}
	fmt.Printf("1. scale= %+v, Abs= %v \n", v, v.Abs())

	v.Scale(5)
	fmt.Printf("2. scale= %+v, Abs= %v \n", v, v.Abs())

}
