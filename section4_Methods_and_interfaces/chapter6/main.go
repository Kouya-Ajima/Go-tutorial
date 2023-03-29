package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// v.Scale の場合、メソッド呼び出し時に勝手にポインター変数が渡される。
	// (&v).Scale と解釈される。
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{3, 4}
	p.Scale(2)
	ScaleFunc(p, 10)

	// v,p のどちらでやっても結果は同じ。
	// ただ、p は、ポインタの参照が返却される
	fmt.Println(v, p)
}
