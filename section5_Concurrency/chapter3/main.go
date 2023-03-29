package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//error
	// ch <- 4

	// 受信と表示を同時に実行
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
