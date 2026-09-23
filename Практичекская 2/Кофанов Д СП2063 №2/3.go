package main

import "fmt"

type Order struct {
	id          int
	items       []int
	total       float64
	address     string
	isCompleted bool
}

var orderMap = make(map[int]Order)

func addOrder(order Order) {
	orderMap[order.id] = order
}

func main() {
	newOrder := Order{
		id:          101,
		items:       []int{2, 11, 23},
		total:       1550.50,
		address:     "г.Курган, ул.Пушкина",
		isCompleted: false,
	}
	addOrder(newOrder)
	fmt.Printf("Добавлен заказ: %+v\n", orderMap[101])
}
