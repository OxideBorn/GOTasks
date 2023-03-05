package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	_, err := fmt.Scanf("%f", &x)
	if err != nil {
		fmt.Println("n/a")
	} else {
		fmt.Printf("%.1f", 7*math.Pow10(-3)*math.Pow(x, 4)+((22.8*math.Pow(x, 1/3)-math.Pow10(3))*x+3)/(x*x/2)-x*math.Pow((10+x), (2/x))-1.01)
	}
}
