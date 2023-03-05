package main

import (
	"fmt"
	"math"
)

var x, y float64

func main() {

	_, err := fmt.Scanf("%f%f", &x, &y)
	if err != nil {
		fmt.Println("n/a")
		fmt.Println(err)

	} else {
		if math.Pow(x, 2)+math.Pow(y, 2) <= 25 {
			fmt.Println("GOTCHA")

		} else {
			fmt.Println("MISS")
		}
	}
}
