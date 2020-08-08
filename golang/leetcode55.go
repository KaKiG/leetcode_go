package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	maxIndex := 0
	for i := range nums {
		if maxIndex >= i {
			maxIndex = Max2(maxIndex, i+nums[i])
		}
		if maxIndex >= len(nums)-1 {
			return true
		}
	}
	return false
}

func Max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// a simple but time consuming way
/*func canJump(nums []int) bool {
if len(nums) == 0 {
	return false
}

state := make([]bool, len(nums))
state[0] = true
for i := 0; i < len(nums); i++ {
	if state[i] == true {
		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			state[j] = true
		}
	}
}
return state[len(nums)-1]
*/
