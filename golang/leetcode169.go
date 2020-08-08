package main

import "fmt"

func main() {
	nums := []int{1, 1, 0, 0, 0, 1, 0}
	fmt.Println(majorityElement(nums))
}

//quicker method, boyer moore voting
//we set a candidate and a count to both zero at first.
//Range over the array, if we see a number same as the candidate
//count++, else count--
//if count == 0, we set candidate to the next number we meet.

func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	for i := range nums {
		if count == 0 {
			candidate = nums[i]
		}

		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// func majorityElement(nums []int) int {
// 	numMap := make(map[int]int, 0)
//
// 	for i := range nums {
// 		numMap[nums[i]]++
// 	}
//
// 	for k, v := range numMap {
// 		if v > len(nums)/2 {
// 			return k
// 		}
// 	}
// 	return -1
// }
