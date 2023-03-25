package main

import (
	"fmt"
	"strings"
)

func main() {

	// [3][各ストリングの配列] を生成
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "x"
	board[2][2] = "O"
	board[1][2] = "x"
	board[1][0] = "O"
	board[0][2] = "x"

	// strings.Join()で、配列の二次元目を全て表示させる
	fmt.Println(strings.Join(board[0], "　　 "))

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
