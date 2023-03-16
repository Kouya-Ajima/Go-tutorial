package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// 上から順にTrueを当てはめてBreakする。
	switch {
	case t.Hour() < 12:
		fmt.Println("moring")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening")
	}
}
