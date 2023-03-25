package main

import (
	"fmt"
)

func adder() func(int) int {
	// 初期化され、変数部に返される値。
	// 変数代入のときだけ実行される。
	sum := 0
	return func(x int) int {
		// ここで変数に値を代入して更新している
		sum += x
		// 標準出力に返す用
		return sum
	}
}

func main() {
	// 変数に、Adderの初期値を格納
	pos, neg := adder(), adder()

	// ここには関数が格納されているだけ
	fmt.Println(pos)
	// 元の変数に加算される
	fmt.Println(pos(10))

	// 初期化
	pos = adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

}
