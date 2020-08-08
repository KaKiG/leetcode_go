package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 6}
	fmt.Println(subsets(nums))
	fmt.Println(len(subsets(nums)))
}

func subsets(nums []int) [][]int {
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
		PartialSolver(append(append([]int{}, nums[:i]...), nums[i+1:]...), i, ans, append(curr, nums[i]))
	}
}
