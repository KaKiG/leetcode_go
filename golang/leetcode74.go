package main

import "fmt"

func main() {
	matrix := [][]int{
		{},
	}
	target := 1
	fmt.Println(searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	up := 0
	down := len(matrix) - 1
	col := len(matrix[0]) - 1
	i := (up + down) / 2
	for up <= down {
		i = (up + down) / 2
		if matrix[i][col] < target {
			up = i + 1
		}
		if matrix[i][0] > target {
			down = i - 1
		}
		if matrix[i][0] <= target && matrix[i][col] >= target {
			return BinarySearch(matrix[i], target)
		}
	}
	return false
}

func BinarySearch(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	i := (left + right) / 2
	if nums[i] == target {
		return true
	}
	for left <= right {
		i = (left + right) / 2
		if nums[i] < target {
			left = i + 1
		}
		if nums[i] > target {
			right = i - 1
		}
		if nums[i] == target {
			return true
		}
	}
	return false
}
