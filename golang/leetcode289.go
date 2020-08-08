package main

import "fmt"

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(board)
	fmt.Println(board)
}

func gameOfLife(board [][]int) {
	tmpBoard := make([][]int, len(board))
	for i := range tmpBoard {
		tmpBoard[i] = make([]int, len(board[i]))
	}

	for i := range board {
		for j := range board[i] {
			if (board[i][j] == 1 && count(board, i, j) == 2) || count(board, i, j) == 3 {
				tmpBoard[i][j] = 1
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = tmpBoard[i][j]
		}
	}
}

func count(board [][]int, i, j int) int {
	count := 0
	row := len(board)
	col := len(board[0])
	for m := i - 1; m <= i+1; m++ {
		for n := j - 1; n <= j+1; n++ {
			if !(m == i && n == j) && inField(row, col, m, n) {
				count += board[m][n]
			}
		}
	}
	return count
}

func inField(row, col, i, j int) bool {
	if i >= 0 && i < row && j >= 0 && j < col {
		return true
	}
	return false
}
