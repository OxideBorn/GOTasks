package main

import (
	"fmt"
)

func main() {

	var Num1, Num2 int32 //Задаем параметры переиенных

	_, err := fmt.Scanf("%d%d", &Num1, &Num2)
	if err != nil {
		fmt.Println("n/a")
	} else {
		if Num2 == 0 {
			fmt.Println(Num1+Num2, Num1-Num2, Num1*Num2, "n/a")
		} else {
			fmt.Println(Num1+Num2, Num1-Num2, Num1*Num2, Num1/Num2)
		}

	}
}
