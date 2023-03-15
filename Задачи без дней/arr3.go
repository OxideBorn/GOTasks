package main

import "fmt"

func main() {
	var (
		mas [2][5]int
		n   int
	)
	fmt.Scan(&n)
	mas[0][n-1] = n
	mas[1][n-1] = n * n
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(mas[i][j], " ")
		}
		fmt.Print("\n")
	}

}
