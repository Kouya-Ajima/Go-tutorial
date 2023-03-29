package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 関数に書き直した場合
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// *v の参照を変更しているので、元の変数も書き換わる。
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
	// &v への参照を引数で渡す
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
