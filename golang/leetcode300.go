package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

// Binary search method
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp := make([]int, 0)
	dp = append(dp, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
		} else {
			index := BinarySearch(dp, nums[i])
			dp[index] = nums[i]
		}
	}
	return len(dp)
}

func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2

	for left < right {
		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

// func lengthOfLIS(nums []int) int {
// 	if len(nums) <= 1 {
// 		return len(nums)
// 	}

// 	dp := make([]int, len(nums))
// 	dp[len(nums)-1] = 1
// 	maxLen := 0

// 	for i := len(nums) - 2; i >= 0; i-- {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[j] > nums[i] {
// 				dp[i] = max2(dp[i], dp[j])
// 			}
// 		}
// 		dp[i]++
// 		maxLen = max2(maxLen, dp[i])
// 	}
// 	// fmt.Println(dp)
// 	return maxLen
// }

// func max2(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
