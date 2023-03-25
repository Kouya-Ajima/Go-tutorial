package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m, k map[string]Vertex

func main() {

	// リテラルで初期化しない場合、Makeが必要
	m = make(map[string]Vertex)

	k = make(map[string]Vertex)

	m["Bell Lads"] = Vertex{
		40.4482, -77.43223,
	}

	k["Bell Lads"] = Vertex{
		20.4482, 37.43223,
	}

	fmt.Println(m["Bell Lads"], k["Bell Lads"])
}
