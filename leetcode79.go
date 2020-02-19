package main

func main() {

}

func exist(board [][]byte, word string) bool {
	var check bool
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				board[i][j] = '.'
				check = PartialSolver(board, i, j, word[1:])
				if check == true {
					return check
				}
				board[i][j] = word[0]
			}
		}
	}
	return false
}

func PartialSolver(board [][]byte, i, j int, word string) bool {
	if word == "" {
		return true
	}
	var check bool
	tmp := word[0]
	if Infield(board, i-1, j) && board[i-1][j] == tmp {
		board[i-1][j] = '.'
		check = check || PartialSolver(board, i-1, j, word[1:])
		board[i-1][j] = tmp
	}
	if Infield(board, i, j+1) && board[i][j+1] == tmp {
		board[i][j+1] = '.'
		check = check || PartialSolver(board, i, j+1, word[1:])
		board[i][j+1] = tmp
	}
	if Infield(board, i+1, j) && board[i+1][j] == tmp {
		board[i+1][j] = '.'
		check = check || PartialSolver(board, i+1, j, word[1:])
		board[i+1][j] = tmp
	}
	if Infield(board, i, j-1) && board[i][j-1] == tmp {
		board[i][j-1] = '.'
		check = check || PartialSolver(board, i, j-1, word[1:])
		board[i][j-1] = tmp
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
