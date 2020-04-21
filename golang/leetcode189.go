package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	if len(nums) == 0 || len(nums) == 1 || k == 0 {
		return
	}

	new := make([]int, len(nums))
	copy(new, nums[len(nums)-k:len(nums)])
	copy(new[k:], nums)
	copy(nums, new)
}
