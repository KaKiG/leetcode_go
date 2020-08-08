package main

import "fmt"

func main() {
	s := "babgbag"
	t := "bag"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s string, t string) int {
	if len(s) == 0 {
		if len(t) == 0 {
			return 1
		}
		return 0
	}

	dpTable := make([][]int, len(s)+1)
	for i := range dpTable {
		dpTable[i] = make([]int, len(t)+1)

	}

	for i := range dpTable {
		dpTable[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dpTable[i][j] = dpTable[i-1][j-1] + dpTable[i-1][j]
			} else {
				dpTable[i][j] = dpTable[i-1][j]
			}
		}
	}
	return dpTable[len(s)][len(t)]
}
