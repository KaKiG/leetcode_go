package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{2, 3},
		{5, 5},
		{2, 2},
		{3, 4},
		{3, 4},
	}
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	if len(intervals) == 1 {
		return intervals
	}
	//sortedIntervals := SortIntervals(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	sortedIntervals := intervals
	res := make([][]int, 0)
	tmp := make([][]int, 0)
	tmp = append(tmp, sortedIntervals[0])
	for i := 0; i < len(sortedIntervals)-1; i++ {
		if len(tmp) == 1 {
			tmp = MergeTwoIntervals(tmp[0], sortedIntervals[i+1])
		} else {
			res = append(res, tmp[0])
			tmp = MergeTwoIntervals(tmp[1], sortedIntervals[i+1])
		}
		if i == len(sortedIntervals)-2 {
			res = append(res, tmp...)
		}
	}
	return res
}

func MergeTwoIntervals(a, b []int) [][]int {
	if a[1] >= b[0] {
		return [][]int{[]int{a[0], Max2(a[1], b[1])}}
	}
	return [][]int{a, b}
}

func SortIntervals(intervals [][]int) [][]int {
	leftIndex := make([]int, len(intervals))
	for i := range intervals {
		leftIndex[i] = intervals[i][0]
	}
	sort.Ints(leftIndex)
	sortedIntervals := make([][]int, len(intervals))
	repeatMark := make([]bool, len(intervals))
	for i := range leftIndex {
		for j := range intervals {
			if repeatMark[j] == false {
				if intervals[j][0] == leftIndex[i] {
					sortedIntervals[i] = intervals[j]
					repeatMark[j] = true
					break
				}
			}
		}
	}
	return sortedIntervals
}

func Max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
