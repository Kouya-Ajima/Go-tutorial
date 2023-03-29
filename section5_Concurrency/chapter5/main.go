package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		// select は、受信準備のできた case を選択して実行する。
		// 両方できてれば、ランダム実行。
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			// ここをチャネル受信にすると、go func() 内の <ーc と競合
			//  → デットロックが発生する。
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		// この関数では、順番に実行される。
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		// for文が終わったら実行。 → ｃの受信準備ができたら実行
		quit <- 0
	}()

	// go func() と一緒に同実行される。
	fibonacci(c, quit)
}
