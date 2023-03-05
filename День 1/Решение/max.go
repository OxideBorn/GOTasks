package main

import (
	"fmt"
)

var s = "n/a"

func main() {
	var num1, num2 int
	_, err := fmt.Scanf("%d %d", &num1, &num2)
	if err != nil {
		fmt.Println(s)
	} else {
		fmt.Println(findMax(num1, num2))
	}

}
func findMax(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
