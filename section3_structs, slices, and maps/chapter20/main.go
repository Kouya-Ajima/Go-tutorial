package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {

	// m = make(map[string]Vertex)

	// リテラルで一緒に初期化しているのでMakeは不要
	m = map[string]Vertex{
		"hoge": Vertex{
			443, 4242,
		},
		"bar": Vertex{
			3232322, 54343,
		},
	}

	fmt.Println(m, m["hoge"].Lat)
}
