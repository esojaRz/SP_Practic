package main

import "fmt"

const (
	single      = "single"
	double      = "double"
	suite       = "suite"
	free        = "free"
	booked      = "booked"
	maintenance = "maintenance"
)

type HotelRoom struct {
	Type   string
	Status string
	Price  float64
}

func bookRoom(rooms map[string]HotelRoom, roomNumber string) {
	room, exists := rooms[roomNumber]
	if exists && room.Status == free {
		room.Status = booked
		rooms[roomNumber] = room
		fmt.Printf("Номер %s забронирован.\n", roomNumber)
	} else {
		fmt.Printf("Бронирование номера %s невозможно.\n", roomNumber)
	}
}
func main() {
	hotel := map[string]HotelRoom{
		"15": {Type: single, Status: free, Price: 2500},
		"16": {Type: double, Status: maintenance, Price: 4000},
	}

	bookRoom(hotel, "15")
}
