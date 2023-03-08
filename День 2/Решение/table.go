package main

import (
	"fmt"
	"time"
)

func main() {
	// true {

	//for a := time.Now().Hour(); a <= 24; a++ {
	//	for b := time.Now().Minute(); b < 60; b++ {
	//		for c := time.Now().Second(); c < 60; c++ {
	//			cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	//			cmd.Stdout = os.Stdout
	//			cmd.Run()
	//			fmt.Printf("Часы: %d Минуты: %d Секунды: %d\n", a, b, c)
	//			time.Sleep(1 * time.Second)
	//		}
	//	}
	//}
	//cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	//cmd.Stdout = os.Stdout
	//cmd.Run

	//}
	var y, m, d int
	fmt.Print("Введите день.месяц.год:")
	_, err := fmt.Scanf("%d.%d.%d", &d, &m, &y)
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Print(y, m, d)
	t1 := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	t2 := time.Now()
	fmt.Println((int(t2.Sub(t1) / (time.Hour * 24))) / 365)

}
