/*
 * @lc app=leetcode id=304 lang=golang
 *
 * [304] Range Sum Query 2D - Immutable
 */

// @lc code=start
type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sumMat := make([][]int, len(matrix))
	for i := range sumMat {
		sumMat[i] = make([]int, len(matrix[0]))
	}
	for i := range sumMat {
		for j := range sumMat[i] {
			if i == 0 && j == 0 {
				sumMat[i][j] = matrix[i][j]
			} else if i == 0 {
				sumMat[i][j] = matrix[i][j] + sumMat[i][j-1]
			} else if j == 0 {
				sumMat[i][j] = matrix[i][j] + sumMat[i-1][j]
			} else {
				sumMat[i][j] = matrix[i][j] + sumMat[i-1][j] + sumMat[i][j-1] - sumMat[i-1][j-1]
			}
		}
	}
	return NumMatrix{sumMat}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	all := this.sumMatrix[row2][col2]
	left := 0
	up := 0
	leftUp := 0
	if row1 != 0 {
		left = this.sumMatrix[row1-1][col2]
	}
	if col1 != 0 {
		up = this.sumMatrix[row2][col1-1]
	}
	if row1 != 0 && col1 != 0 {
		leftUp = this.sumMatrix[row1-1][col1-1]
	}
	return all - left - up + leftUp
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end
