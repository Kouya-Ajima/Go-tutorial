package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	fmt.Println(n)

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// for文が終わったらクローズ → 後にRangeループをクローズする時だけやる
	close(c)
}

func main() {
	c := make(chan int, 10)
	// チャネルのバッファ → キャパシティ として、int で渡せる。
	fibonacci(cap(c), c)

	// この場合の range は、チャネルが終了するまで続ける。
	// fibonacci() で、すでに close(c) で終了される命令が入っている。
	for i := range c {
		// rangeでチャネルをi に入れる → 受信したことになる。
		fmt.Println(i)
	}
}
