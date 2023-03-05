package main

import (
	"fmt"
)

func main() {
	var x, y float64
	_, err := fmt.Scanf("%f%f", &x, &y)

	if err != nil {
		fmt.Println("n/a")

	} else {
		if x*x+y*y <= 25 {
			fmt.Println("GOTCHA")
		} else {
			fmt.Println("MISS")
		}

	}
}
