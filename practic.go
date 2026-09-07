package main

import (
	"fmt"
	"math"
)

func main() {

	//1 Задача
	rent := 95000.0
	NewRent := rent * 1.10
	fmt.Printf("новая стоимость аренды: %.2f руб.\n", NewRent)

	//2 Задача
	laptops := 55480 * 6
	monitors := 21830 * 3
	mice := 890 * 11
	keyboards := 1560 * 5
	sum := laptops + monitors + mice + keyboards
	fmt.Printf("Общая сумма покупки: %d руб.\n", sum)

	//3 Задача
	a := 5000
	b := 256
	c := a / b
	d := a % b
	fmt.Printf("Можно разместить файлов: %d\n", c)
	fmt.Printf("останется свободного места: %d Гб\n", d)

	//4 Задача
	var temp float64
	fmt.Print("Введите температуру в фаренгейтах: ")
	fmt.Scan(&temp)
	cel := 5.0 / 9.0 * (temp - 32)
	fmt.Printf("Температура в цельсиях: %.2f C\n", cel)

	//5 Задача
	var r float64
	fmt.Print("Введите радиус клумбы: ")
	fmt.Scan(&r)
	L := 2 * math.Pi * r
	S := math.Pi * r * r
	fmt.Printf("Длина окружности: %.2f\n", L)
	fmt.Printf("Площадь круга: %.2f\n", S)

	//6 Задача
	var i, a1, years float64
	fmt.Print("Введите начальную сумму вклада: ")
	fmt.Scan(&i)
	fmt.Print("Введите годовую процентную ставку: ")
	fmt.Scan(&a1)
	fmt.Print("Введите количество лет: ")
	fmt.Scan(&years)
	result := i * math.Pow(1+a1/100, years)
	fmt.Printf("Итоговая сумма вклада: %.2f\n", result)

	//9 Задача
	var suma float64
	fmt.Print("Введите сумму покупки: ")
	fmt.Scan(&suma)
	disc := suma * 0.8
	fmt.Printf("Сумма со скидкой 20%%: %.2f\n", disc)

	//10 Задача
	var A, B float64
	fmt.Print("Введите число a: ")
	fmt.Scan(&A)
	fmt.Print("Введите число b: ")
	fmt.Scan(&B)
	delenie := A / B
	roundedUp := math.Ceil((delenie))
	roundedDown := math.Floor(float64(delenie))
	fmt.Printf("Результат деления: %.4f\n", delenie)
	fmt.Printf("Округление вверх: %.0f\n", roundedUp)
	fmt.Printf("Округление вверх: %.0f\n", roundedDown)
}
