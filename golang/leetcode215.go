package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
