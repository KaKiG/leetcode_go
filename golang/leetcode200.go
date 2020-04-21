package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	count := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				count++
				grid[i][j] = '0'
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if i-1 >= 0 && grid[i-1][j] == '1' {
		grid[i-1][j] = '0'
		dfs(grid, i-1, j)
	}

	if i+1 < len(grid) && grid[i+1][j] == '1' {
		grid[i+1][j] = '0'
		dfs(grid, i+1, j)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		grid[i][j-1] = '0'
		dfs(grid, i, j-1)
	}

	if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
		grid[i][j+1] = '0'
		dfs(grid, i, j+1)
	}
}
