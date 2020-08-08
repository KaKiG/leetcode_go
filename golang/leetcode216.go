package main

import "fmt"

func main() {
	k, n := 3, 9
	fmt.Println(combinationSum3(k, n))
}

func combinationSum3(k int, n int) [][]int {
	if n > 45 || k > 9 {
		return nil
	}

	res := make([][]int, 0)
	solver(n, 1, k, &res, nil)
	return res
}

func solver(target, start, length int, res *[][]int, ans []int) {
	if target == 0 && length == 0 {
		*res = append(*res, append([]int{}, ans...))
	}

	for i := start; i <= 9; i++ {
		solver(target-i, i+1, length-1, res, append(ans, i))
	}
}
