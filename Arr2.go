package main

import "fmt"

func main() {

	var arr [5]int
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)
}
