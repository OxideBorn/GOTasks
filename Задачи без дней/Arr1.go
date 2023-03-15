package main

import (
	"fmt"
	"math"
)

func main() {
	var mas [2][5]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			if i == 0 {
				mas[i][j] = j + 1
			} else {
				mas[i][j] = int(math.Pow(float64(2), float64(j+1)))
			}
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(mas[i][j], " ")
		}
		fmt.Print("\n")
	}
}
