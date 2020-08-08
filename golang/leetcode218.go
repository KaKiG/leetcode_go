package main

import (
	"fmt"
)

func main() {
	buildings := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}
	fmt.Println(getSkyline(buildings))
}

func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}

	if len(buildings) == 1 {
		return [][]int{
			{buildings[0][0], buildings[0][2]},
			{buildings[0][1], 0},
		}
	}
	m := len(buildings) / 2
	return merge(getSkyline(buildings[:m]), getSkyline(buildings[m:]))
}

func merge(skyline1, skyline2 [][]int) [][]int {
	res := make([][]int, 0)

	var i, j, x, h1, h2 int
	for i < len(skyline1) && j < len(skyline2) {
		// for each point, find the maximum height. This is possible because two skylines have been sorted
		if skyline1[i][0] == skyline2[j][0] {
			x = skyline1[i][0]
			h1 = skyline1[i][1]
			h2 = skyline2[j][1]
			i++
			j++
		} else if skyline1[i][0] < skyline2[j][0] {
			x = skyline1[i][0]
			h1 = skyline1[i][1]
			i++
		} else {
			x = skyline2[j][0]
			h2 = skyline2[j][1]
			j++
		}
		res = appendSkyline(res, x, max2(h1, h2))
	}

	for ; i < len(skyline1); i++ {
		res = appendSkyline(res, skyline1[i][0], skyline1[i][1])
	}

	for ; j < len(skyline2); j++ {
		res = appendSkyline(res, skyline2[j][0], skyline2[j][1])
	}

	return res
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func appendSkyline(skyline [][]int, x, y int) [][]int {
	// if no new point is made with different height, do not append anything
	if len(skyline) == 0 || skyline[len(skyline)-1][1] != y {
		return append(skyline, []int{x, y})
	}
	return skyline
}
