package main

import "fmt"

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 2, 3}
	fmt.Println(minSubArrayLen(s, nums))
}

func minSubArrayLen(s int, nums []int) int {
	sum := 0
	lens := len(nums)
	left := -1
	right := 0
	for right < len(nums) {
		for right < len(nums) && sum < s {
			sum += nums[right]
			right++
		}
		lens = min2(lens, right-left-1)

		for left < len(nums) && sum >= s {
			left++
			sum -= nums[left]
		}
		lens = min2(lens, right-left)

	}
	if right == len(nums) && left == -1 && sum < s {
		return 0
	}

	return lens
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
