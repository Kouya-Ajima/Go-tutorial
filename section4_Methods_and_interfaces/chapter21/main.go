package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		// Read() は 第一引数に、読み取ったByteの値の文字数を返す。
		//  → 第2引数に、エラーを返す → 何も負ければ Nil
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b =%v \n", n, err, b)

		fmt.Printf("b[:n] = %q \n", b[:n])
		// 読み取りが終わったら EOF のエラーを返す。
		if err == io.EOF {
			break
		}
	}
}
