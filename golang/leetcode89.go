package main

import "fmt"

func main() {
	fmt.Println(grayCode(3))
}

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	prev := grayCode(n - 1)
	m := len(prev)
	for i := 0; i < m; i++ {
		prev = append(prev, prev[m-i-1]+Pow2(n-1))
	}
	return prev
}

func Pow2(n int) int {
	sum := 1
	for i := 0; i < n; i++ {
		sum *= 2
	}
	return sum
}
