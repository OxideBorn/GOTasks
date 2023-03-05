package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int

	_, err := fmt.Scanf("%d%d", &x, &y)
	if err != nil {
		fmt.Println("n/a")
	} else {
		//fmt.Println(math.Max(float64(x), float64(y)))// в сотаве функции main
		fmt.Println(findMax(x, y))
	}
}
func findMax(x, y int) float64 {
	//var m float64 альтернатива
	//m=math.Max(x,y)
	return math.Max(float64(x), float64(y))

}
