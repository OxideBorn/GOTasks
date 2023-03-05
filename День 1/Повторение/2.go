package main

import (
	"fmt"
)

var name int

func main() {
	fmt.Scan(&name)

	fmt.Println("Hello, " + fmt.Sprint(name) + "!")

}
