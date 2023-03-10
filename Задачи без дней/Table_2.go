package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)
	for i := 1; i <= 10; i++ {
		fmt.Println(x, "x", i, "=", x*i)
	}

}
