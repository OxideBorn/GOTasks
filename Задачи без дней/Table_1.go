package main

import (
	"fmt"
)

var x, y int

func main() {

	fmt.Scan(&x, &y)

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			fmt.Print(i*j, "\t")

		}
		fmt.Println()
	}

}
