package main

import "fmt"

type val_ind struct {
	val int
	ind int
}

func main() {
	nums1 := []int{5, 2, 6, 1}
	fmt.Println(countSmaller(nums1))
	nums2 := []int{5, 2, 6, 1, 7}
	fmt.Println(countSmaller(nums2))
	nums3 := []int{1, 2, 3, 4, 543, 66, 3, 45}
	fmt.Println(countSmaller(nums3))
	nums4 := []int{}
	fmt.Println(countSmaller(nums4))
	nums5 := []int{1}
	fmt.Println(countSmaller(nums5))
}

func countSmaller(nums []int) []int {
	val_ind_arr := make([]val_ind, len(nums))
	count_arr := make([]int, len(nums))
	for i := range nums {
		val_ind_arr[i].val = nums[i]
		val_ind_arr[i].ind = i
	}

	if len(count_arr) > 1 {
		mergeSort(val_ind_arr[:len(nums)/2], val_ind_arr[len(nums)/2:], &count_arr)
	}
	return count_arr
}

func mergeSort(left []val_ind, right []val_ind, count_arr *[]int) []val_ind {
	res_val_ind_arr := make([]val_ind, 0)
	sorted_left_ind := 0
	sorted_right_ind := 0
	right_smaller_num := 0
	var sorted_left []val_ind
	var sorted_right []val_ind

	if len(left) == 1 && len(right) == 1 {
		if left[0].val <= right[0].val {
			return []val_ind{left[0], right[0]}
		} else {
			(*count_arr)[left[0].ind]++
			return []val_ind{right[0], left[0]}
		}
	}

	if len(left) == 1 {
		sorted_left = left
	} else {
		sorted_left = mergeSort(left[:len(left)/2], left[len(left)/2:], count_arr)
	}
	sorted_right = mergeSort(right[:len(right)/2], right[len(right)/2:], count_arr)
	for sorted_left_ind != len(sorted_left) || sorted_right_ind != len(sorted_right) {
		if sorted_left[sorted_left_ind].val <= sorted_right[sorted_right_ind].val {
			res_val_ind_arr = append(res_val_ind_arr, sorted_left[sorted_left_ind])
			(*count_arr)[sorted_left[sorted_left_ind].ind] += right_smaller_num
			sorted_left_ind++
		} else {
			res_val_ind_arr = append(res_val_ind_arr, sorted_right[sorted_right_ind])
			sorted_right_ind++
			right_smaller_num++
		}

		if sorted_left_ind == len(sorted_left) && sorted_right_ind != len(sorted_right) {
			res_val_ind_arr = append(res_val_ind_arr, sorted_right[sorted_right_ind:]...)
			break
		}

		if sorted_left_ind != len(sorted_left) && sorted_right_ind == len(sorted_right) {
			res_val_ind_arr = append(res_val_ind_arr, sorted_left[sorted_left_ind:]...)
			for i := sorted_left_ind; i < len(sorted_left); i++ {
				(*count_arr)[sorted_left[i].ind] += right_smaller_num
			}
			break
		}
	}
	return res_val_ind_arr
}
