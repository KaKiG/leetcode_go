package main

import "fmt"

func main() {
	nums := []int{1, 5, 9, 1, 5, 9}
	k, t := 2, 3
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t))
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}

	if t == 0 {
		return containsNearbyDuplicate(nums, k)
	}
	//here we try to scale the numbers down by t
	//therefore, if two numbers are different by at most t, then after transformation
	//they differ by at most 1. Then we can apply the thought we used in 219
	//also we need to deal with the case when number is negative
	numMap := make(map[int][2]int)
	for i, v := range nums {
		if v < 0 {
			v -= t - 1
		}

		if val, exist := numMap[v/t]; exist && i-val[0] <= k {
			return true
		}

		if val, exist := numMap[v/t+1]; exist && i-val[0] <= k && val[1]-v <= t {
			return true
		}

		if val, exist := numMap[v/t-1]; exist && i-val[0] <= k && v-val[1] <= t {
			return true
		}

		numMap[v/t] = [2]int{i, v}
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)

	for i := range nums {
		_, exist := numMap[nums[i]]
		if !exist {
			numMap[nums[i]] = i
		} else {
			if i-numMap[nums[i]] <= k {
				return true
			}
			numMap[nums[i]] = i
		}
	}
	return false
}
