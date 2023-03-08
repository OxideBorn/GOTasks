package main

var number1, number2 int

func main() {
	
}

func isRemains(number1, number2 int) bool {
	for number1 >= number2 {
		number1 -= number2
		//fmt.Println(number1, number2)
	}
	return number1 != 0
}
