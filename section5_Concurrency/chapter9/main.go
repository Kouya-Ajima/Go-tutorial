package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc は、与えられたキーのカウンターをインクリメントする
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// map c.v は
	// 一度に1つのゴルーチンだけがマップにアクセスできるようにロックする
	c.v[key]++
	c.mu.Unlock()
}

// Value は、与えられたキーに対するカウンターの現在値を返す。
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// map c.v は
	// 一度に1つのゴルーチンだけがマップにアクセスできるようにロックする
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
