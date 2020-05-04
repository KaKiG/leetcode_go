package main

import "sort"

func main() {

}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	var diff int
	for i := 0; i < len(nums)-1; i++ {
		diff = max2(diff, nums[i+1]-nums[i])
	}
	return diff
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
