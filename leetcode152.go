package main

import "fmt"

func main() {
	nums := []int{-2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := make([]int, len(nums))
	min := make([]int, len(nums))
	max[0] = nums[0]
	min[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		max[i], min[i] = maxMin3(max[i-1]*nums[i], min[i-1]*nums[i], nums[i])
	}
	return maxArray(max)
}

func maxMin3(a, b, c int) (int, int) {
	return max2(max2(a, b), c), min2(min2(a, b), c)
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArray(nums []int) int {
	max := nums[0]
	for i := range nums {
		max = max2(max, nums[i])
	}
	return max
}
