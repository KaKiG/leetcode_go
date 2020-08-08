package main

import "fmt"

func main() {
	n := 1000
	fmt.Println(countDigitOne(n))
}

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	count := 0
	count += n / 10
	if n%10 != 0 {
		count++
	}
	level := 10

	for level <= n {
		count += n/(level*10)*level + min2(max2(n%(10*level)-level+1, 0), level)
		level *= 10
	}
	return count
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
