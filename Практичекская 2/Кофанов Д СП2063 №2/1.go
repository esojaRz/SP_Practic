package main

import "fmt"

func main() {
	a := 4
	b := 4

	weekdayPrice := 2100
	weekendPrice := 2850
	totalSum := (a * weekdayPrice) + (b * weekendPrice)
	fmt.Println(totalSum)
}
