package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// どちらでも動く。
	t := T{"Hello World"}
	t.M()
	// この時、Implementsは必要ない。
	var i I = T{"  Hello World2"}
	i.M()

}
