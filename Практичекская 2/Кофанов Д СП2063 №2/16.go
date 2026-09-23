package main

import (
	"fmt"
	"time"
)

type Sensor struct {
	ID   string
	Temp float64
	Hum  float64
	Time time.Time
}

func calc(list []Sensor) float64 {
	sum := 0.0
	for _, x := range list {
		sum += x.Temp
	}
	return sum / float64(len(list))
}

func main() {
	t := time.Now()

	arr := []Sensor{
		{"S1", 20.0, 50.0, t},
		{"S1", 22.5, 48.0, t},
		{"S1", 25.0, 45.0, t},
		{"S1", 21.0, 51.0, t},
		{"S1", 19.5, 53.0, t},
	}

	res := calc(arr)
	fmt.Println("Средняя температура:", res)
}
