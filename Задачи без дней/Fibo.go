package main

import "fmt"

func main() {
	//Задать число n
	//y=0+1+2+3+4+5
	var n int
	//var res, res1 int = 0, 1
	var num, num1 = 0, 1

	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		num, num1 = num+num1, num
		//fmt.Print(num+num1, "\t")
	}
	fmt.Printf("%d член числа Фибоначчи,%d", n, num+num1)
}
