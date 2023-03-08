package main

import (
	"fmt"
)

func main() {
	var a int
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Println("n/a")

	} else {
		var x int
		for i := 1; i <= a; i++ {
			if isSimple(i) == true && isDivider(i, a) == true {
				x = i
			}
		}
		fmt.Println(x)
	}
}

func isSimple(i int) bool {
	//var count = 0
	//for j:=1; j<=i; j++ {
	//	if i%j == 0 {
	//		count++
	//	}
	//}
	//if count >2 {
	//	return false
	//}
	//return true
	for j := 1; j <= i; j++ {
		if i%j == 0 && j != 1 && j != i {
			return false
		}
	}
	return true
}

func isDivider(i, a int) bool {
	if a%i == 0 {
		return true
	}
	return false
}
