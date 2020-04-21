package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	maxArea := 0
	stack = append(stack, -1)
	for i := range heights {
		for len(stack) != 1 && heights[i] <= heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxArea = max2(maxArea, heights[top]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}

	for i := 1; i < len(stack); i++ {
		maxArea = max2(maxArea, (len(heights)-stack[i-1]-1)*heights[stack[i]])
	}
	return maxArea
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//try solutio with stack
/*func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	area := make([]int, len(heights))
	for i := range heights {
		area[i] = Area(heights, i)
	}
	return MaxArray(area)
}

func Area(heights []int, i int) int {
	width := 1
	for j := i + 1; j < len(heights); j++ {
		if heights[j] >= heights[i] {
			width++
		} else {
			break
		}
	}
	for j := i - 1; j >= 0; j-- {
		if heights[j] >= heights[i] {
			width++
		} else {
			break
		}
	}
	return width * heights[i]
}
func MaxArray(nums []int) int {
	max := nums[0]
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}*/
