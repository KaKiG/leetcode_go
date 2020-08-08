package main

import (
	"fmt"
)

func main() {
	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}

	fmt.Println(calculateMinimumHP(dungeon))
}

// We can use dynamic programming from bottom right to upper left
// opt represents the least healthy value needed before calculate the current grid
// Therefore, we can derive a recursive formula the compute the least healthy value
// Be sure to know at least 1 is needed for survive
func calculateMinimumHP(dungeon [][]int) int {
	row := len(dungeon)
	col := len(dungeon[0])

	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if i == row-1 && j == col-1 {
				dungeon[i][j] = max2(1, 1-dungeon[i][j])
				continue
			}

			if i == row-1 {
				dungeon[i][j] = max2(1, dungeon[i][j+1]-dungeon[i][j])
				continue
			}

			if j == col-1 {
				dungeon[i][j] = max2(1, dungeon[i+1][j]-dungeon[i][j])
				continue
			}

			dungeon[i][j] = max2(1, min2(dungeon[i+1][j], dungeon[i][j+1])-dungeon[i][j])
		}
	}
	return dungeon[0][0]
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
