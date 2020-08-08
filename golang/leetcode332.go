package main

import (
	"fmt"
	"sort"
)

func main() {
	tickets := [][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}
	fmt.Println(findItinerary(tickets))
}

func findItinerary(tickets [][]string) []string {
	if len(tickets) == 0 {
		return []string{}
	}

	airMap := make(map[string][]string)
	visitedMap := make(map[string][]bool)

	for i := range tickets {
		airMap[tickets[i][0]] = append(airMap[tickets[i][0]], tickets[i][1])
		visitedMap[tickets[i][0]] = append(visitedMap[tickets[i][0]], false)
	}

	for k := range airMap {
		sort.Strings(airMap[k])
	}

	str := "JFK"
	res := []string{}

	dfs([]string{str}, &res, airMap, visitedMap, str, len(tickets))
	return res
}

// here the biggest problem is that dictionary will change the memory directly, it is not safe to do so in dfs
// so we use another dictionary (bool) to record whether a string have been used or not
func dfs(curr []string, res *[]string, airMap map[string][]string, visitedMap map[string][]bool, str string, i int) {
	if i == 0 {
		*res = curr
		return
	}
	for j, v := range airMap[str] {
		if len(*res) != 0 {
			return
		}
		if visitedMap[str][j] == false {
			visitedMap[str][j] = true
			dfs(append(curr, v), res, airMap, visitedMap, v, i-1)
			visitedMap[str][j] = false
		}

	}
}
