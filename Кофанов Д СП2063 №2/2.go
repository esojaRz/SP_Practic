package main

import "fmt"

func main() {
	var mainBag, handLuggage, extraHandLuggage float64

	fmt.Print("Введите вес основного багажа: ")
	fmt.Scan(&mainBag)

	fmt.Print("Введите вес ручной клади: ")
	fmt.Scan(&handLuggage)

	fmt.Print("Введите вес дополнительной ручной клади: ")
	fmt.Scan(&extraHandLuggage)

	totalWeight := mainBag + handLuggage + extraHandLuggage
	fmt.Printf("Общий вес багажа: %.2f кг\n", totalWeight)
}
