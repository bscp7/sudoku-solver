package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Sudoku struct {
	Board     [9][9]string
	Solved    bool
	TimeTaken time.Duration
}

func (s *Sudoku) NewBoardFromFile(filepath string) {
	fmt.Printf("filepath: %v\n", filepath)
	fileReader, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(fileReader)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	board := [9][9]string{}
	for x, line := range lines {
		numbers := strings.Split(line, "")
		for i := 0; i < 9; i++ {
			board[x][i] = numbers[i]
		}
	}
	s.Board = board
	fileReader.Close()
}

func (s *Sudoku) Print() {
	printSudoku(&s.Board)
}

func (s *Sudoku) Solve(delay time.Duration) {
	start := time.Now()
	solveSudoku(&s.Board, delay)
	if isSudokuSolved(&s.Board) {
		s.Solved = true
		s.TimeTaken = time.Since(start)
	}
}

func (s *Sudoku) IsSolved() bool {
	s.Solved = isSudokuSolved(&s.Board)
	return s.Solved
}

func (s *Sudoku) PrintStats() {
	fmt.Printf("Solved: %v\nTimeTaken: %v\n", s.Solved, s.TimeTaken)
}
