package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	recMap := make(map[int]bool)
	res := make([]int, 0)

	for _, v := range nums1 {
		recMap[v] = true
	}

	for _, v := range nums2 {
		if recMap[v] {
			res = append(res, v)
			recMap[v] = false
		}
	}
	return res
}
