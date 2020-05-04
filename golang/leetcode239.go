package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	slide := make([]int, 0)

	for i := range nums {

		if len(slide) > 0 && slide[0] < i-k+1 {
			slide = slide[1:]
		}

		for len(slide) > 0 && nums[slide[len(slide)-1]] < nums[i] {
			slide = slide[:len(slide)-1]
		}

		slide = append(slide, i)

		if i >= k-1 {
			res[i-k+1] = nums[slide[0]]
		}
	}

	return res
}
