package main

import (
	"strconv"
	"strings"
	"time"
)

func finalNumbersInRow(board *[9][9]string, row int) string {
	var numbersInRow string
	for i := 0; i < len(board); i++ {
		var placeValue = board[row][i]
		if placeValue != "0" && len(placeValue) == 1 {
			numbersInRow += placeValue
		}
	}
	return numbersInRow
}

func finalNumbersInCol(board *[9][9]string, col int) string {
	var numbersInCol string
	for i := 0; i < len(board); i++ {
		var placeValue = board[i][col]
		if placeValue != "0" && len(placeValue) == 1 {
			numbersInCol += placeValue
		}
	}
	return numbersInCol
}

func finalNumbesInBlock(board *[9][9]string, row int, col int) string {
	var numbersInBlock string
	if row <= 2 {
		row = 0
	} else if row >= 3 && row <= 5 {
		row = 3
	} else {
		row = 6
	}
	if col <= 2 {
		col = 0
	} else if col >= 3 && col <= 5 {
		col = 3
	} else {
		col = 6
	}
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			var placeValue = board[i][j]
			if placeValue != "0" && len(placeValue) == 1 {
				numbersInBlock += board[i][j]
			}
		}
	}
	return numbersInBlock
}

func removeDuplicatesFromString(value string) string {
	var noDuplicate string
	numbers := strings.Split(value, "")
	for i := 0; i < len(numbers); i++ {
		if !strings.Contains(noDuplicate, numbers[i]) {
			noDuplicate += numbers[i]
		}
	}
	return noDuplicate
}

func missingNumbersInString(value string) string {
	var missing string
	for i := 1; i <= 9; i++ {
		if !strings.Contains(value, strconv.Itoa(i)) {
			missing += strconv.Itoa(i)
		}
	}
	return missing
}

func finalNumbers(board *[9][9]string, row int, col int) string {
	fnRow := finalNumbersInRow(board, row)
	fnCol := finalNumbersInCol(board, col)
	fnBlk := finalNumbesInBlock(board, row, col)
	fn := removeDuplicatesFromString(fnRow + fnCol + fnBlk)
	return fn
}

func fillPredictionsInEmptyBlocks(board *[9][9]string, delay time.Duration) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			time.Sleep(delay * time.Millisecond)
			placeValue := board[i][j]
			if placeValue == "0" {
				fn := finalNumbers(board, i, j)
				// using dot "." as an indicator of predicted values
				board[i][j] = "." + missingNumbersInString(fn)
			}
		}
	}
}

func removePredictedNumberFromRow(board *[9][9]string, row int, numberToBeRemoved string) {
	for i := 0; i < 9; i++ {
		placeValue := board[row][i]
		if strings.HasPrefix(placeValue, ".") && strings.Contains(placeValue, numberToBeRemoved) {
			placeValue = strings.Join(strings.Split(placeValue, numberToBeRemoved), "")
			board[row][i] = strings.Replace(placeValue, ".", "_", -1)
		}
	}
}

func removePredictedNumberFromCol(board *[9][9]string, col int, numberToBeRemoved string) {
	for j := 0; j < 9; j++ {
		placeValue := board[j][col]
		if strings.HasPrefix(placeValue, ".") && strings.Contains(placeValue, numberToBeRemoved) {
			placeValue = strings.Join(strings.Split(placeValue, numberToBeRemoved), "")
			board[j][col] = strings.Replace(placeValue, ".", "_", -1)
		}
	}
}

func removePredictedNumberFromBlock(board *[9][9]string, row int, col int, numberToBeRemoved string) {
	if row <= 2 {
		row = 0
	} else if row >= 3 && row <= 5 {
		row = 3
	} else {
		row = 6
	}
	if col <= 2 {
		col = 0
	} else if col >= 3 && col <= 5 {
		col = 3
	} else {
		col = 6
	}
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			var placeValue = board[i][j]
			if strings.HasPrefix(placeValue, ".") && strings.Contains(placeValue, numberToBeRemoved) {
				placeValue = strings.Join(strings.Split(placeValue, numberToBeRemoved), "")
				board[i][j] = strings.Replace(placeValue, ".", "_", -1)
			}
		}
	}
}

func cleanupBoard(board *[9][9]string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			placeValue := board[i][j]
			if strings.HasPrefix(placeValue, "_") {
				placeValue = strings.Join(strings.Split(placeValue, "_"), "")
				board[i][j] = "." + placeValue
			}
		}
	}
}

func stickSinglePredictionToPlace(board *[9][9]string, delay time.Duration) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			time.Sleep(delay * time.Millisecond)
			// remove dot "." from string and check if length of the string is 1
			// if it is then set the number to a place
			placeValue := board[i][j]
			if strings.HasPrefix(placeValue, ".") {
				placeValue = strings.Join(strings.Split(placeValue, "."), "")
			}
			if len(placeValue) == 1 {
				// fix the placevalue
				board[i][j] = placeValue
				// when you do add new place value,
				// ensure that the number is removed
				// from other other predicated values in same row, col
				removePredictedNumberFromRow(board, i, placeValue)
				removePredictedNumberFromCol(board, j, placeValue)
				removePredictedNumberFromBlock(board, i, j, placeValue)
			}
		}
	}
}

func isSudokuSolved(board *[9][9]string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if len(board[i][j]) > 1 {
				return false
			}
		}
	}
	return true
}

func solveSudoku(board *[9][9]string, delay time.Duration) {
	fillPredictionsInEmptyBlocks(board, delay)

	for {
		if isSudokuSolved(board) {
			break
		}
		stickSinglePredictionToPlace(board, delay)
		cleanupBoard(board)
	}
}
