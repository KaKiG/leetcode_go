package main

import (
	"fmt"
)

func main() {
	envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	fmt.Println(maxEnvelopes(envelopes))
}

// brute force, time limit exceeds
func maxEnvelopes(envelopes [][]int) int {
	ans := 0
	solver(envelopes, 0, 0, 0, &ans)
	return ans
}

func solver(enve [][]int, w, h, curr int, ans *int) {
	if curr > *ans {
		*ans = curr
	}
	for i := range enve {
		if enve[i][0] > w && enve[i][1] > h {
			solver(append(append([][]int{}, enve[:i]...), enve[i+1:]...), enve[i][0], enve[i][1], curr+1, ans)
		}
	}
}
