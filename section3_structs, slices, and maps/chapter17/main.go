package main

import (
	"fmt"
)

func main() {

	pow := make([]int, 10)

	// 配列に2の階乗を格納
	// 1 << uint(i) の１は基数。 ２ にすると２の階乗になる。
	// 基数１ に、値（インデックス）の２乗をかけ合わせたのも
	for i := range pow {
		pow[i] = 1 << uint(i) // 2のi乗
	}

	fmt.Println(pow)

	for _, v := range pow {
		fmt.Printf("%d \n", v)
	}
}
