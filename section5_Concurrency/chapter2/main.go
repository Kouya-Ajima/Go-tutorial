package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	// 第１がインデックス、第2が値
	for _, v := range s {
		sum += v
	}
	// sum をチャネルに送信
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	// Mapのように、チャネルを生成する
	c := make(chan int)

	// ゴルーチンで使用しないとデットロックが発生。
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)

	// チャネルに送信した変数を受診して格納
	// チャネルは、必ず受信する必要がある。
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
