package main

import "fmt"

func main() {
	m, n := 6, 7
	fmt.Println(rangeBitwiseAnd(m, n))
}

// naive solution, but O(n)
// there is a O(1) method
func rangeBitwiseAnd(m int, n int) int {
	b := uint(0)
	k := n - m
	for k != 0 {
		k = k >> 1
		b++
	}
	return (n >> b << b) & (m >> b << b)
}
