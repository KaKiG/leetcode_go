package main

import "fmt"

func main() {
	board := [][]byte{
		{'a', 'b'},
		{'a', 'a'},
	}
	words := []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}
	fmt.Println(findWords(board, words))
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	for _, v := range words {
		newBoard := make([][]bool, len(board))
		for i := range newBoard {
			newBoard[i] = make([]bool, len(board[0]))
		}

		if exist(board, newBoard, v) {
			res = append(res, v)
		}
	}
	return res
}

func exist(board [][]byte, newBoard [][]bool, word string) bool {
	var check bool
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				newBoard[i][j] = true
				check = PartialSolver(board, newBoard, i, j, word[1:])
				if check == true {
					return check
				}
				newBoard[i][j] = false
			}
		}
	}
	return false
}

func PartialSolver(board [][]byte, newBoard [][]bool, i, j int, word string) bool {
	if word == "" {
		return true
	}
	var check bool
	if Infield(board, i-1, j) && board[i-1][j] == word[0] && newBoard[i-1][j] == false {
		newBoard[i-1][j] = true
		check = check || PartialSolver(board, newBoard, i-1, j, word[1:])
		newBoard[i-1][j] = false
	}
	if Infield(board, i, j+1) && board[i][j+1] == word[0] && newBoard[i][j+1] == false {
		newBoard[i][j+1] = true
		check = check || PartialSolver(board, newBoard, i, j+1, word[1:])
		newBoard[i][j+1] = false
	}
	if Infield(board, i+1, j) && board[i+1][j] == word[0] && newBoard[i+1][j] == false {
		newBoard[i+1][j] = true
		check = check || PartialSolver(board, newBoard, i+1, j, word[1:])
		newBoard[i+1][j] = false
	}
	if Infield(board, i, j-1) && board[i][j-1] == word[0] && newBoard[i][j-1] == false {
		newBoard[i][j-1] = true
		check = check || PartialSolver(board, newBoard, i, j-1, word[1:])
		newBoard[i][j-1] = false
	}
	return check
}

func Infield(board [][]byte, i, j int) bool {
	row := len(board)
	col := len(board[0])

	if i >= 0 && i < row && j >= 0 && j < col {
		return true
	}
	return false
}
