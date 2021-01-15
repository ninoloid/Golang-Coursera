package main

import (
	"fmt"
	"math"
)

func trunc(x float64) float64 {
	return math.Trunc(x)
}

func main() {
	var userInput float64
	fmt.Scan(&userInput)
	fmt.Println(trunc(userInput))
}
