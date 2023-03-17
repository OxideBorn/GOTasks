package main

import "fmt"

func main() {

	var n, u, m, s int

	fmt.Scanf("%d%d%d%d", &n, &u, &m, &s)

	//fmt.Scan(&n, &u, &m, &s)

	res := (n + u + m + s)

	fmt.Println(res)

}
