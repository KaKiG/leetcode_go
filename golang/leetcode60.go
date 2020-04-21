package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	n := 4
	k := 9
	fmt.Println(getPermutation(n, k))
}

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	for i := 1; i < k; i++ {
		nextPermutation(nums)
	}
	str := ""
	for i := range nums {
		str += strconv.Itoa(nums[i])
	}
	return str
}

func nextPermutation(nums []int) {
	index := IndexDeOrder(nums)
	if index == -1 {
		sort.Ints(nums)
	} else {
		indexNum, restSlice := NextGNumRest(nums[index+1:], nums[index])
		nums = append(nums[:index], indexNum)
		nums = append(nums, restSlice...)
	}
}

func IndexDeOrder(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			return i - 1
		}
	}
	return -1
}

func NextGNumRest(nums []int, a int) (int, []int) {
	temp := append(nums, a)
	index := -1
	sort.Ints(temp)
	for i := range temp {
		if temp[i] > a {
			index = i
			break
		}
	}
	if index == len(nums) {
		resInt := temp[index]
		res := temp[:index]
		return resInt, res
	}
	resInt := temp[index]
	res := append(temp[:index], temp[index+1:]...)
	return resInt, res
}
