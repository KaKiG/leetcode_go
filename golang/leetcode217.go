package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	numDic := make(map[int]int)

	for i := range nums {
		if numDic[nums[i]] == 1 {
			return true
		}
		numDic[nums[i]]++
	}
	return false
}
