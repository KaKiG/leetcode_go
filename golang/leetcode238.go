package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

// Using two passes
// The first one calculates the multiplication of 0 to i-1 for all num[i]
// The second one calculates the multiplication of n to i+1 for all num[i]
// Multiply them together
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	sum := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		ans[i] *= sum
		sum *= nums[i]
	}
	return ans
}
