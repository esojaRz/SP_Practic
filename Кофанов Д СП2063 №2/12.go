package main

import (
	"fmt"
	"strconv"
)

const (
	baseBin = 2
	baseDec = 10
	baseHex = 16
)

func convertNumber(valueStr string, sourceBase int, targetBase int) (string, error) {
	decimalValue, err := strconv.ParseInt(valueStr, sourceBase, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(decimalValue, targetBase), nil
}

func main() {
	var inputNumber int64 = 1654
	binaryRepresentation := strconv.FormatInt(inputNumber, baseBin)
	hexRepresentation := strconv.FormatInt(inputNumber, baseHex)
	fmt.Printf("Исходное число (десятичное): %d\n", inputNumber)
	fmt.Printf("В двоичной системе: %s\n", binaryRepresentation)
	fmt.Printf("В шестнадцатеричной системе: %s\n\n", hexRepresentation)

	hexInput := "ea"
	convertedResult, err := convertNumber(hexInput, baseHex, baseBin)
	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
	} else {
		fmt.Printf("Результат перевода из 16-ричной '%s' в двоичную: %s\n", hexInput, convertedResult)
	}
}
