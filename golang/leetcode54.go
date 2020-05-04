package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	right := len(matrix[0]) - 1
	left := 0
	top := 0
	down := len(matrix) - 1
	res := make([]int, 0)

	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		if top > down {
			break
		}

		for i := top; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if top > down {
			break
		}

		for i := down; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
