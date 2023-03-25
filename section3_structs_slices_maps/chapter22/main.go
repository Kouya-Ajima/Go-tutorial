package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["Answer"] = 32324
	fmt.Println(m["Answer"])

	m["Answer"] = 55
	fmt.Println(m["Answer"])

	// キューの順序、FIFOで処理される。
	delete(m, "Amswer")
	fmt.Println(m["Answer"])

	// キーがあるかの確認は、下記のように第２引数に格納される
	// 同じキーがある場合、キュー形式なので、追加分が参照される
	v, isOk := m["Answer"]
	fmt.Println(v, isOk)
}
