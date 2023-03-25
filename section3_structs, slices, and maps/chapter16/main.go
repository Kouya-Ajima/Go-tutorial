package main

import (
	"fmt"
)

func main() {

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	// Rangeは、インデックスと値をそれぞれ返す、Mapのときとかにも使える
	for i, v := range pow {
		fmt.Printf("   i = %d,  v = %d \n", i, v)
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
