package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	max := 1
	curr := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			continue
		}
		if nums[i+1] == nums[i]+1 {
			curr++
		} else {
			max = max2(max, curr)
			curr = 1
		}
	}

	return max2(curr, max)
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
