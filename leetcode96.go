package main

import "fmt"

func main() {
	fmt.Println(numTrees(5))
}

type TreeKey struct {
	start, end int
}

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	TMap := make(map[TreeKey]int)
	return Solver(1, n, TMap)
}

func Solver(start, end int, TMap map[TreeKey]int) int {
	if start > end {
		return 1
	}

	if start == end {
		return 1
	}
	if _, exist := TMap[TreeKey{start, end}]; exist {
		return TMap[TreeKey{start, end}]
	}

	var ans int
	for i := start; i <= end; i++ {
		ans += Solver(start, i-1, TMap) * Solver(i+1, end, TMap)
	}

	TMap[TreeKey{start, end}] = ans
	return ans
}
