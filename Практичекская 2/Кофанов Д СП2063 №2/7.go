package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

func calculateSalary(employees []Employee) (float64, float64) {
	if len(employees) == 0 {
		return 0, 0
	}
	total := 0.0
	for _, emp := range employees {
		total += emp.Salary
	}
	return total, total / float64(len(employees))
}

func main() {
	staff := []Employee{
		{ID: 1, Name: "Даня", Position: "Разработчик", Salary: 140000},
		{ID: 2, Name: "Темка", Position: "Темщик", Salary: 155555},
	}

	totalFunds, avgSalary := calculateSalary(staff)
	fmt.Printf("Общий фонд: %.2f, Средняя зарплата: %.2f\n", totalFunds, avgSalary)
}
