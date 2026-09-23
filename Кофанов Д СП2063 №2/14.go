package main

import "fmt"

type backpackItem struct {
	Title       string
	Mass        float64
	IsEssential bool
}

func calculateTotalMass(bag []backpackItem) float64 {
	var total float64
	for _, object := range bag {
		total += object.Mass
	}
	return total
}

func main() {
	loot := []backpackItem{
		{Title: "Меч", Mass: 15.4, IsEssential: false},
		{Title: "Факел", Mass: 0.5, IsEssential: false},
		{Title: "Ключ", Mass: 0.1, IsEssential: true},
		{Title: "Железный щит", Mass: 22.0, IsEssential: false},
	}

	weightSum := calculateTotalMass(loot)
	fmt.Printf("Полный вес снаряжения героя: %.2f кг\n", weightSum)
}
