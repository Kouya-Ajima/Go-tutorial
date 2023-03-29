package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		// select → goroutineを複数の通信操作で待機させる。
		// なにも当てはまらなければDefaultが呼ばれる。
		select {
		case <-tick:
			fmt.Println("tick. ")
		case <-boom:
			fmt.Println("BOOM ! ")
			return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
