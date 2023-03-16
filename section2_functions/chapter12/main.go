package main

import "fmt"

func main() {
	// 関数が終わったタイミングで実行される
	defer fmt.Println("world")

	fmt.Println("Hello")
}
