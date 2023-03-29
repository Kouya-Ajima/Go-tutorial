package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

// メソッドを構造体やタイプに紐付ける -> クラスを表現する。
// 関数の前に Struct,Type をつけることで、クラスのメソッドとして扱える
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.y + v.x*v.y)
}

// 引数もOK
func (v Vertex) Abs2(k float64) float64 {
	return math.Sqrt(v.x*v.y + v.x*k)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v.Abs2(5))
}
