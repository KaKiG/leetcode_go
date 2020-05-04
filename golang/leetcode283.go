package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 0, 0, 3, 4, 5}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	posn := 0
	for i := range nums {
		if nums[i] != 0 {
			if i != posn {
				nums[posn], nums[i] = nums[i], 0
			}
			posn++
		}
	}
}
