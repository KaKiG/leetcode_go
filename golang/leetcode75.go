package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 0, 1, 2}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	if len(nums) > 0 {
		count0 := 0
		count1 := 0
		count2 := 0
		for _, v := range nums {
			if v == 0 {
				count0++
			}
			if v == 1 {
				count1++
			}
			if v == 2 {
				count2++
			}
		}
		for i := 0; i < len(nums); i++ {
			if count0 > 0 {
				nums[i] = 0
				count0--
			} else if count1 > 0 {
				nums[i] = 1
				count1--
			} else {
				nums[i] = 2
				count2--
			}
		}
	}
}
