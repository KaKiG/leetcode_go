package main

import "fmt"

func main() {
	numCourses := 3
	prerequisites := [][]int{
		{1, 0},
		{2, 1},
	}
	fmt.Println(canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	edge := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	queue := make([]int, 0)

	for i := range prerequisites {
		edge[prerequisites[i][1]] = append(edge[prerequisites[i][1]], prerequisites[i][0])
		indegree[prerequisites[i][0]]++
	}

	for i := range indegree {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		numCourses--
		for i := range edge[curr] {
			indegree[edge[curr][i]]--
			if indegree[edge[curr][i]] == 0 {
				queue = append(queue, edge[curr][i])
			}
		}
	}
	return numCourses == 0

}

// too slow
// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	if len(prerequisites) == 0 {
// 		return true
// 	}

// 	numMap := make(map[int]int)
// 	row := len(prerequisites) - 1
// 	col := len(prerequisites[0]) - 1
// 	currPos := make([]int, row+1)
// 	sum := 0
// 	label := 0

// 	for i := range currPos {
// 		currPos[i] = col
// 	}
// 	for i := range prerequisites {
// 		for j := range prerequisites[i] {
// 			numMap[prerequisites[i][j]]++
// 		}
// 	}

// 	for sum != (row+1)*(col+1) && label != numCourses {
// 		calMap := make(map[int]int)
// 		for i := range prerequisites {
// 			if currPos[i] >= 0 {
// 				calMap[prerequisites[i][currPos[i]]]++
// 			}
// 		}

// 		labelMap := make(map[int]bool)
// 		for i := range prerequisites {
// 			if currPos[i] >= 0 && numMap[prerequisites[i][currPos[i]]] == calMap[prerequisites[i][currPos[i]]] {
// 				labelMap[prerequisites[i][currPos[i]]] = true
// 				currPos[i]--
// 				sum++
// 			}
// 		}

// 		if len(labelMap) == 0 {
// 			return false
// 		}
// 		label += len(labelMap)

// 	}
// 	return true

// }
