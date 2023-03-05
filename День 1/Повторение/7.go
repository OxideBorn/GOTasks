package main

import "fmt"

func main() {

	var x int

	_, err := fmt.Scanf("%d", &x)
	if err != nil || x < 0 {
		fmt.Println("n/a")
	} else {
		for i := 1; i <= x; i++ {
			if i%3 == 0 && i%5 == 0 {
				fmt.Print("FIZZBUZZ ")
			} else if i%5 == 0 {
				fmt.Print("BUZZ ")
			} else if i%3 == 0 {
				fmt.Print("FIZZ ")
			} else {
				fmt.Printf("%d ", i)
			}
		}
	}
}
