package main

import (
	"fmt"
	"strings"
)

func printLine() {
	fmt.Println(string("\033[31m"), "------------------------------------", string("\033[0m"))
}

func printSudoku(board *[9][9]string) {
	printLine()
	for i := 0; i < len(board); i++ {
		fmt.Print(string("\033[31m"), "|", string("\033[0m"))
		for j := 0; j < len(board[i]); j++ {
			placeValue := strings.Replace(board[i][j], "0", "", -1)
			fmt.Printf(" %1s ", placeValue)
			fmt.Print(string("\033[31m"), "|", string("\033[0m"))
		}
		fmt.Println()
		printLine()
	}
}
