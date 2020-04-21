package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 2, 1, 1}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) []int {
	//the first part is to select the possible candidates
	nums1 := 0
	nums2 := 1
	count1 := 0
	count2 := 0
	for i := range nums {
		if nums[i] == nums1 {
			count1++
		} else if nums[i] == nums2 {
			count2++
		} else if count1 <= 0 {
			count1, nums1 = 1, nums[i]
		} else if count2 <= 0 {
			count2, nums2 = 1, nums[i]
		} else {
			count1--
			count2--
		}
	}

	count1 = 0
	count2 = 0
	//the second part is to check whether they are truely valid
	for i := range nums {
		if nums[i] == nums1 {
			count1++
		}
		if nums[i] == nums2 {
			count2++
		}

	}

	num := []int{}
	if count1 > len(nums)/3 {
		num = append(num, nums1)
	}
	if count2 > len(nums)/3 {
		num = append(num, nums2)
	}
	return num
}
