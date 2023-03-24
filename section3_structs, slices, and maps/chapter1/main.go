package main

import (
	"fmt"
)

func main() {

	i, j := 42, 2701

	// ｉへのポインター
	p := &i
	// ｐの参照先
	fmt.Println(*p)

	*p = 21
	fmt.Println(*p)

	p = &j
	fmt.Println(p)
	fmt.Println(*p)

	*p = *p / 37
	fmt.Println(*p)
	fmt.Println(j)

}
