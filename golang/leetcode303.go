package main

import "fmt"

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	fmt.Println(obj.dp)
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}

type NumArray struct {
	dp [][]int
}

func Constructor(nums []int) NumArray {
	var arr NumArray
	arr.dp = make([][]int, len(nums))
	for i := range arr.dp {
		arr.dp[i] = make([]int, len(nums))
	}

	for i := range nums {
		arr.dp[i][i] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			arr.dp[i][j] = arr.dp[i][j-1] + nums[j]
		}
	}

	return arr
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.dp[i][j]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
