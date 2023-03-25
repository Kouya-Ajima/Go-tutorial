package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {

	// ここでの、リテラル初期化 ＋ 宣言 でも動く
	var m = map[string]Vertex{
		"hoge": Vertex{
			443, 4242,
		},
		"bar": Vertex{
			3232322, 54343,
		},
	}

	fmt.Println(m, m["hoge"].Lat)
}
