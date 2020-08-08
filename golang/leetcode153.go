package main

import "fmt"

func main() {
	nums := []int{3, 1}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	left, right, mid := 0, len(nums)-1, len(nums)/2

	for left < right {
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return nums[mid]
}
