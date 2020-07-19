/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	count := 0
	countArr := make([][]int, len(grid))
	for i := range countArr {
		countArr[i] = make([]int, len(grid[0]))
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				countArr[i][j] += 4
				if (i+1) < len(grid) && grid[i+1][j] == 1 {
					countArr[i][j]--
					countArr[i+1][j]--
				}
				if (j+1) < len(grid[0]) && grid[i][j+1] == 1 {
					countArr[i][j]--
					countArr[i][j+1]--
				}
				count += countArr[i][j]
			}
		}
	}
	return count
}

func num1(grid [][]int, i, j int) int {
	count := 4
	if (i-1) >= 0 && grid[i-1][j] == 1 {
		count--
	}
	if (i+1) < len(grid) && grid[i+1][j] == 1 {
		count--
	}
	if (j-1) >= 0 && grid[i][j-1] == 1 {
		count--
	}
	if (j+1) < len(grid[0]) && grid[i][j+1] == 1 {
		count--
	}
	return count
}

// @lc code=end
