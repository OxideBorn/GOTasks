package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var s = "n/a"
	_, err := fmt.Scanf("%f", &x)
	if err != nil {
		fmt.Println(s)
	} else {

		y := 7*math.Pow10(-3)*math.Pow(x, 4) + ((22.8*math.Pow(x, 1/3)-math.Pow10(3))*x+3)/(x*x/2) - x*math.Pow((10+x), (2/x)) - 1.01
		fmt.Printf("%.1f", y)
	}
}
