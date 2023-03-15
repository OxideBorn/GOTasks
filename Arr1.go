package main

import "fmt"

func main() {

	fmt.Printf("Введите размер вашего массива: ")

	var size int

	fmt.Scanln(&size)

	var arr = make([]string, size)

	for i := 0; i < size; i++ {

		fmt.Printf("Введите %d: ", i)

		fmt.Scanf("%s", &arr[i])
	}

	fmt.Println("Ваш массив: ", arr)

}
