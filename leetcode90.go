package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return [][]int{}
	}
	ans := make([][]int, 0)
	curr := make([]int, 0)
	PartialSolver(nums, 0, &ans, curr)
	return ans
}

func PartialSolver(nums []int, start int, ans *[][]int, curr []int) {
	*ans = append(*ans, append([]int{}, curr...))

	for i := start; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		PartialSolver(append(append([]int{}, nums[:i]...), nums[i+1:]...), i, ans, append(curr, nums[i]))
	}
}
