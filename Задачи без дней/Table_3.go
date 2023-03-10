package main

import "fmt"

func main() {

	fmt.Println("Число")

	var num1 int

	fmt.Scan(&num1)
	fmt.Println("Cтепень")

	var num2 int

	fmt.Scan(&num2)

	result := num1

	for i := 1; i < num2; i++ {

		result = result * num1
	}

	fmt.Println("Итог\n", result)
}
