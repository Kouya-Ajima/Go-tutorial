package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "This is a pen. This is a piano."
	words := strings.Fields(s)
	fmt.Printf("%q\n", words)
}
