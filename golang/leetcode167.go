package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}

//quicker method
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] < target {
			left++
		}
		if numbers[left]+numbers[right] > target {
			right--
		}
	}
	return []int{}
}

// func twoSum(numbers []int, target int) []int {
// 	numMap := make(map[int]int)
// 	for i := range numbers {
// 		numMap[numbers[i]] = i
// 	}
//
// 	for i := range numbers {
// 		_, exist := numMap[target-numbers[i]]
// 		if exist {
// 			return []int{i + 1, numMap[target-numbers[i]] + 1}
// 		}
// 	}
// 	return []int{}
// }
