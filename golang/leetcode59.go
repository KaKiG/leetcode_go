package main

import "fmt"

func main() {
	n := 0
	fmt.Println(generateMatrix(n))
}

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	right := n - 1
	left := 0
	top := 0
	down := n - 1
	j := 1

	for {
		for i := left; i <= right; i++ {
			matrix[top][i] = j
			j++
		}
		top++
		if top > down {
			break
		}

		for i := top; i <= down; i++ {
			matrix[i][right] = j
			j++
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			matrix[down][i] = j
			j++
		}
		down--
		if top > down {
			break
		}

		for i := down; i >= top; i-- {
			matrix[i][left] = j
			j++
		}
		left++
		if left > right {
			break
		}
	}
	return matrix
}
