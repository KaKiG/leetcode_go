package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rowMax := len(matrix) - 1
	colMax := len(matrix[0]) - 1
	row := 0
	col := colMax

	for row >= 0 && col >= 0 && row <= rowMax && col <= colMax {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else if matrix[row][col] < target {
			row++
		}
	}
	return false
}
