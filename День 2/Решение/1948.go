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
		var maxnumber int
		for i := 1; i <= a; i++ {
			if isSimple(i) == true && isDivider(i, a) == true {
				maxnumber = i
			}
		}
		fmt.Println(maxnumber)
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
		if isRemain(i, j) == 0 && j != 1 && j != i {
			return false
		}
	}
	return true
}

func isDivider(i, a int) bool {
	if isRemain(a, i) == 0 {
		return true
	}
	return false
}
func isRemain(num1, num2 int) int {

	for {
		if num1 >= num2 {
			num1 -= num2

		} else {
			break
		}

	}
	return num1
}
