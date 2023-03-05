package main

import (
	"fmt"
)

func main() {
	var s = "n/a"
	var LeftNumber, RightNumber int16
	_, err := fmt.Scanf("%d%d", &LeftNumber, &RightNumber)
	if err != nil {
		fmt.Println(s)
	} else {
		if RightNumber == 0 {
			fmt.Println(RightNumber+LeftNumber, LeftNumber-RightNumber,
				RightNumber*LeftNumber, fmt.Sprint(s))
		} else {
			fmt.Println(RightNumber+LeftNumber, LeftNumber-RightNumber,
				RightNumber*LeftNumber, RightNumber/LeftNumber)
		}

	}
}
