package main

import "fmt"

func main() {
	financeBook := make(map[string]float64)
	financeBook["Магазины"] = 15000.0
	financeBook["Развлечения"] = 5000.0
	financeBook["Аптека"] = 3000.0

	financeBook["Магазины"] += 2000.0
	var grandTotal float64
	fmt.Println("Отчет по ежемесячным тратам:")

	for group, amount := range financeBook {
		fmt.Printf("- %s: %.2f руб.\n", group, amount)
		grandTotal += amount
	}

	fmt.Printf("\nРасход по всем позициям: %.2f руб.\n", grandTotal)
}
