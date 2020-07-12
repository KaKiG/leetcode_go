package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums1)
	fmt.Println(nums1)
	nums2 := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums2)
	fmt.Println(nums2)
	nums3 := []int{1, 3, 2, 2, 3, 1, 6}
	wiggleSort(nums3)
	fmt.Println(nums3)
}

func wiggleSort(nums []int) {
	if len(nums) > 1 {
		sort.Ints(nums)
		res_arr := make([]int, 0)
		mid := len(nums)/2 + len(nums)%2
		for i := 0; i < mid; i++ {
			res_arr = append(res_arr, nums[mid-i-1])
			if mid+i < len(nums) {
				res_arr = append(res_arr, nums[len(nums)-1-i])
			}
		}

		for i := 0; i < len(nums); i++ {
			nums[i] = res_arr[i]
		}
	}
}
