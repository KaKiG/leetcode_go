package main

import "fmt"

func main() {
	nums := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	// nums := [][]int{
	// 	{7, 8, 9},
	// 	{9, 7, 6},
	// 	{7, 2, 3},
	// }
	// nums := [][]int{{1, 2}}
	fmt.Println(longestIncreasingPath(nums))
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	row := len(matrix)
	col := len(matrix[0])

	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}

	maxVal := 1
	for i := range dp {
		for j := range dp[i] {
			dfs(&dp, matrix, row, col, i, j, &maxVal)
		}
	}
	return maxVal
}

func dfs(dp *[][]int, matrix [][]int, row, col, i, j int, maxVal *int) {
	// here do dfs before take max2 is to make sure the value is optimal
	// and if there is a value, when do not need to do dfs again
	if (*dp)[i][j] != 0 {
		return
	}
	(*dp)[i][j] = 1
	if inBoard(row, col, i-1, j) {
		if matrix[i-1][j] > matrix[i][j] {
			dfs(dp, matrix, row, col, i-1, j, maxVal)
			(*dp)[i][j] = max2((*dp)[i-1][j]+1, (*dp)[i][j])
			*maxVal = max2(*maxVal, (*dp)[i][j])
		}
	}
	if inBoard(row, col, i+1, j) {
		if matrix[i+1][j] > matrix[i][j] {
			dfs(dp, matrix, row, col, i+1, j, maxVal)
			(*dp)[i][j] = max2((*dp)[i+1][j]+1, (*dp)[i][j])
			*maxVal = max2(*maxVal, (*dp)[i][j])
		}
	}
	if inBoard(row, col, i, j-1) {
		if matrix[i][j-1] > matrix[i][j] {
			dfs(dp, matrix, row, col, i, j-1, maxVal)
			(*dp)[i][j] = max2((*dp)[i][j-1]+1, (*dp)[i][j])
			*maxVal = max2(*maxVal, (*dp)[i][j])
		}
	}
	if inBoard(row, col, i, j+1) {
		if matrix[i][j+1] > matrix[i][j] {
			dfs(dp, matrix, row, col, i, j+1, maxVal)
			(*dp)[i][j] = max2((*dp)[i][j+1]+1, (*dp)[i][j])
			*maxVal = max2(*maxVal, (*dp)[i][j])
		}
	}
}

func inBoard(row, col, i, j int) bool {
	if i >= 0 && i < row && j >= 0 && j < col {
		return true
	}
	return false
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
