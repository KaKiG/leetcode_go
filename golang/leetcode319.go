package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(bulbSwitch(10))
}

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
