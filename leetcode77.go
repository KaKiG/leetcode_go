package main

import "fmt"

func main() {
	fmt.Println(combine(5, 2))
}

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	curr := make([]int, 0)
	PartialSolver(n, k, 1, &ans, curr)
	return ans
}

func PartialSolver(n, k, start int, ans *[][]int, curr []int) {
	if len(curr) == k {
		*ans = append(*ans, append([]int{}, curr...))
	}

	for i := start; i <= n; i++ {
		PartialSolver(n, k, i+1, ans, append(curr, i))
	}
}
