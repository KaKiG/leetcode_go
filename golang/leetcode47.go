package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 2}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	ans := make([][]int, 0)
	sort.Ints(nums)
	BacktrackPer(nums, nil, &ans)
	return ans
}

func BacktrackPer(nums []int, prev []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int{}, prev...))
	}

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		BacktrackPer(append(append([]int{}, nums[0:i]...), nums[i+1:]...), append(prev, nums[i]), ans)
	}
}
