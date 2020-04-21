package main

import "fmt"

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	distM := make([][]int, len(triangle))
	for i := range triangle {
		distM[i] = make([]int, len(triangle[i]))
		if i == 0 {
			distM[0][0] = triangle[0][0]
			continue
		}

		distM[i][0] = distM[i-1][0] + triangle[i][0]
		distM[i][len(distM[i])-1] = distM[i-1][len(distM[i-1])-1] + triangle[i][len(triangle[i])-1]
		for j := 1; j < len(distM[i])-1; j++ {
			distM[i][j] = min2(distM[i-1][j], distM[i-1][j-1]) + triangle[i][j]
		}
	}
	return minArray(distM[len(distM)-1])
}

func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minArray(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}
