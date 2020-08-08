package main

import (
	"fmt"
	"sort"
)

func main() {
	citations := []int{1, 1}
	fmt.Println(hIndex(citations))

}

// might improve by binary search after sorting
// but it does not affect overall asympototic running time
// which is exactly problem 275....
func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}

	sort.Ints(citations)
	for i := len(citations); i >= 0; i-- {
		if i == len(citations) && citations[len(citations)-i] >= i {
			return i
		} else if i == 0 {
			return 0
		} else if citations[len(citations)-i] >= i && citations[len(citations)-1-i] <= i {
			return i
		}
	}
	return 0
}
