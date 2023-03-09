package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y, z float64
	fmt.Scanf("%f%f", &x, &y)
	z = math.Sqrt(x*x + y*y)
	fmt.Println(x * x)
	fmt.Println(y * y)
	fmt.Println(z)
}
