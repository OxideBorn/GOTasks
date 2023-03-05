package main

import "fmt"

var name int16

func main() {
	fmt.Scan(&name)
	fmt.Printf("Hello, %d!\n  ", name)
	fmt.Println("Hello, " + fmt.Sprint(name) + "!")
}
