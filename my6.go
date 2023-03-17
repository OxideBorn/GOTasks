package main

import (
	"fmt"

	"math"
)

func main() {

	var y, x float64

	fmt.Scan(&x)

	if x > 1 {
		y = x
		fmt.Println(y)

	} else if x > 0 && x <= 1 {
		y = math.Pow(x, 4)

		fmt.Println(y)
	} else {
		y = math.Pow(x, 2)
		fmt.Println(y)
	}

}
