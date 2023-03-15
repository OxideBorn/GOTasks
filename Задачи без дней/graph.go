package main

func main() {
	var mas [101][11]string
	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			mas[i][j] = "0"
		}
	}
	for x := 1; x <= 10; x++ {
		mas[step(x)][x] = "1"
	}
	for i := 100; i >= 0; i-- {
		for j := 0; j < 10; j++ {
			print(mas[i][j], " ")
		}
		print("\n")
	}
}
func step(x int) int {
	y := x * x
	return y
}
