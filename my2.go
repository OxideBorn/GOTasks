package main

import "fmt"

func main() {

	var n, u, m, s, num1, num2 int

	fmt.Scan(&n, &u, &m, &s)

	if n > u && m > s && u > m {
		num1 = n + u
		num2 = m + s
		fmt.Println(num1, num2)
	} else if s > m && u > n && m > u {
		num1 = s + m
		num2 = n + u
		fmt.Println(num1, num2)
	} else if u > n && m > s && m > n {
		num1 = m + u
		num2 = n + s
		fmt.Println(num1, num2)
	} else {
		num1 = s + m
		num2 = n + u
	}
}
