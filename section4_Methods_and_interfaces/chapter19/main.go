package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// Run() の結果（MyErrorのポインタ）を err に格納して、Nilを調べる
	// Errorの場合、Nilに値が入るので、 err != Nil とかく。
	// run() の実行結果は、errorが返される。このとき、Error()が実行される
	// → それを、関数Error()に オーバーライドすることで、例外処理をする。
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
