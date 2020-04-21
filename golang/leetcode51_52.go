package main

import (
  "fmt"
  "strings"
  )

func main() {
  ans := solveNQueens(5)
  for i := range ans {
    fmt.Println(ans[i])
  }
}

func solveNQueens(n int) [][]string {
	initialStr := InitializeStr(n)
	ans := make([][]string, 0)
	PartialSolver(0, n, initialStr, &ans)

	return ans
}

func PartialSolver(start, n int, currStr [][]string, ans *[][]string) bool {
	if start == n {
		*ans = append(*ans, Condense(currStr))
    // if cap is not enough, go will allocate another memory to store the element appended to it (e.g. double the cap).
    // Thus it will not modify the value outside the function.
    // So we need to use pointer to point to the memory.
    // However, if cap is enough, value outside the function will be modified directly.
		return true
	}
	for i := range currStr[start] {
    if start == 0 {
      currStr = InitializeStr(n)
    }
		currStr[start][i] = "Q"
		if IsValid(currStr) {
			if PartialSolver(start+1, n, currStr, ans) {
        if i != n-1 {
          currStr[start][i] = "."
          continue
        } else {
          currStr[start][i] = "."
          return true
        }
			}
		}
		currStr[start][i] = "."
	}
	return false
}

func InitializeStr(n int) [][]string {
	initalStr := make([][]string, n)
	for i := range initalStr {
		initalStr[i] = make([]string, n)
	}

	for i := range initalStr {
		for j := range initalStr[i] {
			initalStr[i][j] = "."
		}
	}
	return initalStr
}

func IsValid(currStr [][]string) bool {
	xCor := make([]int, 0)
	yCor := make([]int, 0)
	for i := range currStr {
		for j := range currStr[i] {
			if currStr[i][j] == "Q" {
				xCor = append(xCor, i)
				yCor = append(yCor, j)
			}
		}
	}
	if IsDuplicate(xCor) || IsDuplicate(yCor) || IsDiag(xCor, yCor) {
		return false
	}
	return true
}

func IsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func IsDiag(x, y []int) bool {
	if len(x) < 2 {
		return false
	}
	for i := range x {
		for j := i + 1; j < len(x); j++ {
			if AbsInt(x[j]-x[i]) == AbsInt(y[j]-y[i]) {
				return true
			}
		}
	}
	return false
}

func AbsInt(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func Condense(currStr [][]string) []string {
  res := make([]string, len(currStr))
  for i := range res {
    res[i] = strings.Join(currStr[i], "")
  }
  return res
}
