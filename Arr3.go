package main

import "fmt"

func main() {

	var mas [2][2]int

	var i, j, n int

	fmt.Scan(&n)
	j = 0
	i = 0
	fmt.Println(i, j)
	mas[i][j] = 2
	i = n - 2
	j = 2 * i
	mas[i][j] = 4

	fmt.Println(mas)

}
