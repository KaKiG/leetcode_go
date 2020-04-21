package main

import "fmt"

func main() {
	gas := []int{5}
	cost := []int{3}
	fmt.Println(canCompleteCircuit(gas, cost))

}

func canCompleteCircuit(gas []int, cost []int) int {
	cirGas := append(gas, gas...)
	cirCost := append(cost, cost...)
	check := false
	if len(gas) == 1 {
		if gas[0] >= cost[0] {
			return 0
		}
		return -1
	}
	for i := range gas {
		if i == 0 {
			if gas[0] >= cost[0] && gas[len(gas)-1] < cost[len(gas)-1] {
				check = canComplete(cirGas, cirCost, 0)
				if check {
					return 0
				}
			}
		} else {
			if gas[i] >= cost[i] && gas[i-1] < cost[i-1] {
				check = canComplete(cirGas, cirCost, i)
				if check {
					return i
				}
			}
		}
	}
	return -1
}

func canComplete(cirGas []int, cirCost []int, start int) bool {
	store := 0
	for i := start; i < start+len(cirGas)/2; i++ {
		store += cirGas[i] - cirCost[i]
		if store < 0 {
			return false
		}
	}
	return true
}
