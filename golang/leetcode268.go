package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	sum := len(nums)
	for i := range nums {
		sum += i - nums[i]
	}
	return sum
}
