package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	res := []string{}
	start := 0
	for i := 0; i < len(nums)-1; i++ {

		if nums[i+1] != nums[i]+1 {
			res = append(res, convertStr(nums[start], nums[i]))
			start = i + 1
		}
	}

	res = append(res, convertStr(nums[start], nums[len(nums)-1]))
	return res
}

func convertStr(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	}
	return strconv.Itoa(start) + "->" + strconv.Itoa(end)
}
