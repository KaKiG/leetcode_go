package main

import "fmt"

func main() {
	nums := []int{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	dpArray := make([]int, len(nums))
	dpArray[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dpArray[i] = Max2(dpArray[i-1]+nums[i], nums[i])
	}
	return MaxArray(dpArray)
}

func Max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxArray(nums []int) int {
	if len(nums) == 0 {
		return nums[0]
	}

	tmpMax := nums[0]
	for _, v := range nums {
		tmpMax = Max2(v, tmpMax)
	}
	return tmpMax
}

/* Beautiful algorithms to solve the problems, but you have to come up with it first.
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		panic("Invalid input")
	}

	if len(nums) == 1 {
		return nums[0]
	}

	tmpMax := nums[0]
	sum := nums[0]

	for _, n := range nums[1:] {
		if sum > 0 {
			sum += n
		} else {
			sum = n
		}
		tmpMax = Max2(tmpMax, sum)
	}
	return tmpMax
}
*/
