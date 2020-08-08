package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]byte{
		{'0', '1', '1', '0', '1'},
		{'1', '1', '1', '1', '1'},
		{'0', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}

// standard dp solution
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	row := len(matrix)
	col := len(matrix[0])
	dp := make([][]int, row)
	maxSize := 0

	for i := range dp {
		dp[i] = make([]int, col)
	}

	for i := range dp {
		for j := range dp[i] {
			if matrix[i][j] == '0' {
				continue
			}

			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1

			}
			if dp[i][j] > maxSize {
				maxSize = dp[i][j]
			}
		}

	}
	return maxSize * maxSize
}

func min(nums ...int) int {
	minNum := math.MaxInt32

	for i := range nums {
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	return minNum
}

// this idea is modified from MaximumRectangle Area, which can be regarded as more general case
// func maximalSquare(matrix [][]byte) int {
// 	if len(matrix) == 0 {
// 		return 0
// 	}
// 	maxArea := 0
// 	matrixHeight := make([][]int, len(matrix))
// 	for i := range matrixHeight {
// 		matrixHeight[i] = make([]int, len(matrix[0]))
// 	}

// 	for i := range matrixHeight[0] {
// 		if matrix[0][i] == '1' {
// 			matrixHeight[0][i] = 1
// 		} else {
// 			matrixHeight[0][i] = 0
// 		}
// 	}

// 	for i := 1; i < len(matrix); i++ {
// 		for j := range matrixHeight[i] {
// 			if matrix[i][j] == '0' {
// 				matrixHeight[i][j] = 0
// 			} else {
// 				matrixHeight[i][j] = matrixHeight[i-1][j] + 1
// 			}
// 		}
// 	}

// 	for i := range matrix {
// 		maxArea = max2(maxArea, largestSquareArea(matrixHeight[i]))
// 	}
// 	return maxArea
// }

// func largestSquareArea(heights []int) int {
// 	if len(heights) == 0 {
// 		return 0
// 	}
// 	stack := make([]int, 0)
// 	maxArea := 0
// 	stack = append(stack, -1)
// 	for i := range heights {
// 		for len(stack) != 1 && heights[i] <= heights[stack[len(stack)-1]] {
// 			top := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			minLen := min2(heights[top], (i - stack[len(stack)-1] - 1))
// 			maxArea = max2(maxArea, minLen*minLen)
// 		}
// 		stack = append(stack, i)
// 	}

// 	for i := 1; i < len(stack); i++ {
// 		minLen := min2((len(heights) - stack[i-1] - 1), heights[stack[i]])
// 		maxArea = max2(maxArea, minLen*minLen)
// 	}
// 	return maxArea
// }

// func max2(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min2(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
// func minArray(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	min := nums[0]
// 	for i := range nums {
// 		if nums[i] < min {
// 			min = nums[i]
// 		}
// 	}
// 	return min
// }
