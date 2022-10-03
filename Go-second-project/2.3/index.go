package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputString string
	inputString = GetAndConvert("Enter three-digit number:")
	fmt.Println(inputString) //234

	result := ReverseString(inputString)
	fmt.Println(result) // 432
}

func GetAndConvert(numberTitle string) string {
	var strNumber string
	fmt.Print(numberTitle)
	fmt.Scan(&strNumber)
	fmt.Println()
	if len([]rune(strNumber)) > 3 {
		fmt.Println("Invalid Number , number should be three-digit")
		return GetAndConvert(numberTitle)
	}
	_, err := strconv.ParseInt(strNumber, 10, 64)
	if err != nil {
		fmt.Println("Invalid Number")
		return GetAndConvert(numberTitle)
	}
	return strNumber
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
