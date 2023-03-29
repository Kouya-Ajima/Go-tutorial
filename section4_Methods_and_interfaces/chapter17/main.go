package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	// Plintln のときに String() メソッドが呼ばれる。
	// → オーバーライドして使える。
	fmt.Println(a, z)
	fmt.Println("nise", 42)
}
