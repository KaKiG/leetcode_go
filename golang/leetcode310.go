package main

import "fmt"

func main() {
	n := 6
	edges := [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}
	// n := 4
	// edges := [][]int{{1, 0}, {1, 2}, {1, 3}}
	fmt.Println(findMinHeightTrees(n, edges))
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make(map[int][]int)

	for i := range edges {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}

	leaves := make([]int, 0)

	for k, v := range graph {
		if len(v) == 1 {
			leaves = append(leaves, k)
		}
	}

	for n > 2 {
		n -= len(leaves)
		tmpLeaves := make([]int, 0)
		for i := range leaves {
			parents := graph[leaves[i]]

			for _, v := range parents {
				graph[v] = del(graph[v], leaves[i])
				delete(graph, leaves[i])
				if len(graph[v]) == 1 {
					tmpLeaves = append(tmpLeaves, v)
				}
			}
		}

		leaves = []int{}
		leaves = tmpLeaves
	}
	return leaves
}

func del(nums []int, val int) []int {
	for i := range nums {
		if nums[i] == val {
			return append(nums[:i], nums[i+1:]...)
		}
	}
	return nums
}

// too slow, switch dictionary to slice
// func findMinHeightTrees(n int, edges [][]int) []int {
// 	if n == 1 {
// 		return []int{0}
// 	}

// 	graph := make(map[int]map[int]bool)

// 	for i := range edges {
// 		graph[edges[i][0]] = make(map[int]bool)
// 		graph[edges[i][1]] = make(map[int]bool)
// 	}

// 	for i := range edges {
// 		graph[edges[i][0]][edges[i][1]] = true
// 		graph[edges[i][1]][edges[i][0]] = true
// 	}

// 	leaves := make([]int, 0)

// 	for k, v := range graph {
// 		if len(v) == 1 {
// 			leaves = append(leaves, k)
// 		}
// 	}

// 	for n > 2 {
// 		n -= len(leaves)
// 		tmpLeaves := make([]int, 0)
// 		for i := range leaves {
// 			delete(graph, leaves[i])
// 			for _, v := range graph {
// 				delete(v, leaves[i])
// 			}
// 		}

// 		for k, v := range graph {
// 			if len(v) == 1 {
// 				tmpLeaves = append(tmpLeaves, k)
// 			}
// 		}
// 		leaves = []int{}
// 		leaves = tmpLeaves
// 	}

// 	leaves = []int{}
// 	for k := range graph {
// 		leaves = append(leaves, k)
// 	}
// 	return leaves
// }
