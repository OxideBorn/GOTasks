package main

import "fmt"

func main() {
	var n int
	result := 1
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		// rezult:=(n-1)*(n-2)*...(n-(n-1))
		result = result * i
		fmt.Printf("%d\t", result)
	}
}
