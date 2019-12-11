package main

import "fmt"

func main() {
	nums := []int{4, 1, 1, 1, 2, 2, 3, 3}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	numMap := make(map[int]int)

	for i := range nums {
		numMap[nums[i]]++
	}

	for k, v := range numMap {
		if v == 1 {
			return k
		}
	}
	return -1
}
