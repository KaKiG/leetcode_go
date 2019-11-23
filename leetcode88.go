package main

import "fmt"

func main() {

	nums1 := []int{2, 2, 2, 0, 0, 0}
	m := 3
	nums2 := []int{1, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := 0
	index2 := 0
	count1 := 0
	for count1 < m && index2 < n {
		if nums1[index1] <= nums2[index2] {
			index1++
			count1++
		} else {
			//nums[index1:m-count1+index1] = nums[index1+1:m-count1+1+index1]
			for j := m - count1 + index1; j > index1; j-- {
				nums1[j] = nums1[j-1]
			}
			nums1[index1] = nums2[index2]
			index1++
			index2++
		}
	}
	if count1 == m {
		for i := index2; i < n; i++ {
			nums1[index1] = nums2[i]
			index1++
		}
	}
}
