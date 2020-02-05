package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{33, 330}
	// test := []string{"2", "303"}
	fmt.Println(largestNumber(nums))
}

// can be modified using sort.slice become quicker
// replace slice of string
func largestNumber(nums []int) string {
	for i := range nums {
		if nums[i] != 0 {
			str := []string{}
			for i := range nums {
				str = append(str, strconv.Itoa(nums[i]))
			}

			sort.Slice(str, func(i int, j int) bool {
				return str[i]+str[j] > str[j]+str[i]
			})

			return strings.Join(str, "")
		}
	}
	return "0"
}
