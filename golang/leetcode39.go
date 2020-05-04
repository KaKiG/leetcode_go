package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 7, 2}
	target := 18
	fmt.Println(combinationSum(nums, target))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	ComSolve(candidates, target, 0, &res, []int{})
	return res
}

func ComSolve(candidates []int, target int, position int, res *[][]int, ans []int) {
	if target == 0 {
		//this will cause some unexpected modifications to ans in some test cases
		//*res = append(*res, ans)
		//this works and it actually create another space for ans, thus will not disrupte ans.
		/*
		   tmp := make([]int, len(ans))
		   copy(tmp,ans)
		   *res = append(*res, tmp)
		*/
		//some statment as above, just write everything in a line
		*res = append(*res, append([]int{}, ans...))
		return
	}

	for i := position; i < len(candidates); i++ {
		if target-candidates[i] >= 0 {
			ComSolve(candidates, target-candidates[i], i, res, append(ans, candidates[i]))
		}
	}
}
