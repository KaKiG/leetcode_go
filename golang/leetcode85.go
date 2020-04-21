package main

import "fmt"

func main() {
	matrix := [][]byte{
		{'0', '1', '1', '0', '1'},
		{'1', '1', '0', '1', '0'},
		{'0', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '1'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	maxArea := 0
	matrixHeight := make([][]int, len(matrix))
	for i := range matrixHeight {
		matrixHeight[i] = make([]int, len(matrix[0]))
	}

	for i := range matrixHeight[0] {
		if matrix[0][i] == '1' {
			matrixHeight[0][i] = 1
		} else {
			matrixHeight[0][i] = 0
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := range matrixHeight[i] {
			if matrix[i][j] == '0' {
				matrixHeight[i][j] = 0
			} else {
				matrixHeight[i][j] = matrixHeight[i-1][j] + 1
			}
		}
	}

	for i := range matrix {
		maxArea = max2(maxArea, largestRectangleArea(matrixHeight[i]))
	}
	return maxArea
}

func Area(matrixHeight [][]int, i, j int) int {
	width := 1
	for k := j + 1; k < len(matrixHeight[0]); k++ {
		if matrixHeight[i][k] < matrixHeight[i][j] {
			break
		}
		width++
	}
	for k := j - 1; k >= 0; k-- {
		if matrixHeight[i][k] < matrixHeight[i][j] {
			break
		}
		width++
	}
	return width * matrixHeight[i][j]
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	maxArea := 0
	stack = append(stack, -1)
	for i := range heights {
		for len(stack) != 1 && heights[i] <= heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxArea = max2(maxArea, heights[top]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}

	for i := 1; i < len(stack); i++ {
		maxArea = max2(maxArea, (len(heights)-stack[i-1]-1)*heights[stack[i]])
	}
	return maxArea
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
