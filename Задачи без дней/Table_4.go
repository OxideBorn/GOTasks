package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Scan(&num)
	for y := 1; y <= num; y++ {
		fmt.Println(math.Pow(float64(num), float64(y)))
	}
}

func Pow(num, y int) int {
	return int(math.Pow(float64(num), float64(y)))
}
