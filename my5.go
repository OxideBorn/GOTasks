package main

import "fmt"

func main() {

	var y, x float32

	fmt.Scan(&x)
	if x <= 0 {
		y = x
		fmt.Println(y)
	} else if x > 0 && x <= 1 {
		y = x * x
		fmt.Println(y)
	} else {
		y = x * x * x
		fmt.Println(y)

	}
}
