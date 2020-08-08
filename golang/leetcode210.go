package main

import "fmt"

func main() {
	numCourses := 5
	prerequisites := [][]int{}
	fmt.Println(findOrder(numCourses, prerequisites))

}

func findOrder(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0)
	if len(prerequisites) == 0 {
		for i := numCourses - 1; i >= 0; i-- {
			res = append(res, i)
		}
		return res
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
		res = append(res, curr)
		numCourses--
		for i := range edge[curr] {
			indegree[edge[curr][i]]--
			if indegree[edge[curr][i]] == 0 {
				queue = append(queue, edge[curr][i])
			}
		}
	}

	if numCourses == 0 {
		return res
	} else {
		return nil
	}
}
