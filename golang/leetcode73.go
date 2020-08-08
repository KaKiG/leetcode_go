package main

import "fmt"

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := range rows {
		if rows[i] == true {
			SetRowZero(matrix, i)
		}
	}
	for i := range cols {
		if cols[i] == true {
			SetColZero(matrix, i)
		}
	}
}

func SetRowZero(matrix [][]int, i int) {
	for j := range matrix[i] {
		matrix[i][j] = 0
	}
}

func SetColZero(matrix [][]int, j int) {
	for i := range matrix {
		matrix[i][j] = 0
	}
}
