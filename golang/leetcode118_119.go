package main

import "fmt"

func main() {
	fmt.Println(generate(5))
	fmt.Println(getRow(5))
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{
			{1},
		}
	}
	prev := generate(numRows - 1)
	prevRow := prev[len(prev)-1]
	newRow := make([]int, numRows)
	newRow[0], newRow[len(newRow)-1] = 1, 1
	for i := 1; i < len(newRow)-1; i++ {
		newRow[i] = prevRow[i-1] + prevRow[i]
	}
	return append(prev, newRow)
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	prevRow := getRow(rowIndex - 1)
	newRow := make([]int, rowIndex+1)
	newRow[0], newRow[len(newRow)-1] = 1, 1
	for i := 1; i < len(newRow)-1; i++ {
		newRow[i] = prevRow[i-1] + prevRow[i]
	}
	return newRow
}
