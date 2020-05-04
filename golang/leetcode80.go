package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(nums[0:removeDuplicates(nums)])
}

func removeDuplicates(nums []int) int {
	left := 0
	for i := 1; i < len(nums); i++ {
		if left == 0 || nums[i] != nums[left] || nums[left] != nums[left-1] {
			nums[left+1] = nums[i]
			left++
		}
	}
	return left + 1
}
