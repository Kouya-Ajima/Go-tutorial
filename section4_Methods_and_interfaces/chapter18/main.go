package main

import (
	"fmt"
	"reflect"
)

type IPAddr [4]byte

func (ip *IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func print(verb rune, arg interface{}) {
	switch verb {
	case 'v':
		switch v := arg.(type) {
		case fmt.Stringer:
			fmt.Printf("%v: %v\n", reflect.TypeOf(arg).Kind(), v.String())
		default:
			fmt.Printf("%v: %v\n", reflect.TypeOf(arg).Kind(), v)
		}
	}
}

func main() {
	ip := []IPAddr{
		{127, 0, 0, 1},
		{198, 0, 0, 1},
	}
	print('v', ip[0])
	print('v', &ip[1])
}
