package main

import "fmt"

func main() {
	board := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
	}
	solve(board)
	fmt.Println(board)
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	for i := range board[0] {
		if board[0][i] == 'O' {
			bfs(board, 0, i)
		}
		if board[len(board)-1][i] == 'O' {
			bfs(board, len(board)-1, i)
		}
	}

	for i := range board {
		if board[i][0] == 'O' {
			bfs(board, i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			bfs(board, i, len(board[0])-1)
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'Z' {
				board[i][j] = 'O'
			}
		}
	}
}

func bfs(board [][]byte, i, j int) {
	board[i][j] = 'Z'
	if i > 0 && board[i-1][j] == 'O' {
		bfs(board, i-1, j)
	}

	if i < len(board)-1 && board[i+1][j] == 'O' {
		bfs(board, i+1, j)
	}

	if j > 0 && board[i][j-1] == 'O' {
		bfs(board, i, j-1)
	}

	if j < len(board[0])-1 && board[i][j+1] == 'O' {
		bfs(board, i, j+1)
	}
}
