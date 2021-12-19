package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	file            string
	multiPuzzleFile bool
	delay           int
)

func main() {
	flag.StringVar(&file, "f", "", "sudoku puzzle file (9 X 9)")
	flag.BoolVar(&multiPuzzleFile, "m", false, "multiple sudoku puzzles (81 X any)")
	flag.IntVar(&delay, "d", 0, "add delay in processing")
	flag.Parse()

	sudoku := Sudoku{}
	if multiPuzzleFile {
		fileReader, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(fileReader)
		fileScanner.Split(bufio.ScanLines)
		var lines []string
		for fileScanner.Scan() {
			lines = append(lines, fileScanner.Text())
		}
		for _, line := range lines {
			board := [9][9]string{}
			start := 0
			row := 0
			for {
				numbers := strings.Split(line, "")
				for i := 0; i < 9; i++ {
					board[row][i] = numbers[start]
					start += 1
				}
				row += 1

				if start == 81 {
					break
				}
			}

			sudoku.Board = board
			sudoku.Print()
			sudoku.PrintStats()
			sudoku.Solve(time.Duration(delay))
			sudoku.Print()
			sudoku.PrintStats()
		}

		fileReader.Close()
	} else {
		sudoku.NewBoardFromFile(file)
		sudoku.Print()
		sudoku.PrintStats()
		sudoku.Solve(time.Duration(delay))
		sudoku.Print()
		sudoku.PrintStats()
	}
}
