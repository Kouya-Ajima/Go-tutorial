package main

import "fmt"

const Pi = 3.14

func main() {
	// const は 動的な定数
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
