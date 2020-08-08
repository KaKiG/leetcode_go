package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)

	for i := range nums {
		_, exist := numMap[nums[i]]
		if !exist {
			numMap[nums[i]] = i
		} else {
			if i-numMap[nums[i]] <= k {
				return true
			}
			numMap[nums[i]] = i
		}
	}
	return false
}
