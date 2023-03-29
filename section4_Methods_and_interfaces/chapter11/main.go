package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (t *T) M() {
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T) \n", i, i)
}

func main() {
	var i, i2 I
	// 構造体のメッソドを実装するインターフェースに代入するので、 ＆T となる
	i = &T{"Hello World"}
	i2 = F(math.Pi)

	describe(i)
	i.M()
	describe(i2)
	i2.M()
}
