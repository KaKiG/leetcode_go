package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	recMap := make(map[int]int)
	res := make([]int, 0)

	for _, v := range nums1 {
		recMap[v]++
	}

	for _, v := range nums2 {
		if recMap[v] > 0 {
			res = append(res, v)
			recMap[v]--
		}
	}
	return res
}
