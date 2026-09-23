package main

import (
	"errors"
	"fmt"
	"strings"
)

func validateUser(name string, age int, email string) error {
	if name == "" || len(name) >= 50 {
		return errors.New("имя не должно быть пустым или длиннее 50 символов")
	}
	if age < 18 || age > 120 {
		return errors.New("возраст должен быть от 18 до 120 лет")
	}
	if strings.Index(email, "@") == -1 {
		return errors.New("email должен содержать символ '@'")
	}
	return nil
}

func main() {
	err := validateUser("Даня", 20, "danya@gmail.ru")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Данные пользователя валидны.")
	}
}
