package main

import "fmt"

func main() {
	var (
		mas  [2][5]int
		flag = 0
		n    int
	)
	for true {
		fmt.Println("Введите число")
		fmt.Scan(&n)
		for i := 0; i < 5; i++ {
			if mas[0][i] == n {
				fmt.Println("Это число уже введено")
				flag = 1
			}
		}
		if flag == 0 {
			mas[0][n-1] = n
			mas[1][n-1] = n * n
			for i := 0; i < 2; i++ {
				for j := 0; j < 5; j++ {
					fmt.Print(mas[i][j], " ")
				}
				fmt.Print("\n")
			}
		}
		flag = 0
	}
}
