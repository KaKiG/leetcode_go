import "strconv"

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// /*
//  * @lc app=leetcode id=165 lang=golang
//  *
//  * [165] Compare Version Numbers
//  */

// func main() {
// 	version1 := "1.0001"
// 	version2 := "1.1"
// 	fmt.Println(compareVersion(version1, version2))
// }

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	ind1 := 0
	ind2 := 0
	num1, num2 := 0, 0
	for ind1 < len(version1) || ind2 < len(version2) {
		num1, ind1 = findNext(version1, ind1)
		num2, ind2 = findNext(version2, ind2)
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
	}
	return 0
}

func findNext(version string, start int) (int, int) {
	if start > len(version)-1 {
		return 0, start
	}
	i := start
	for i < len(version) {
		if version[i] != '.' {
			i++
		} else {
			break
		}
	}
	versionNum, _ := strconv.Atoi(version[start:i])
	return versionNum, i + 1
}

// @lc code=end
