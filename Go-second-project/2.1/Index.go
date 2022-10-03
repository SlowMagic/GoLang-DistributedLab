// Index.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNumber int64
	firstNumber = GetAndConvertToInt("Enter first numbe: r")
	fmt.Println(firstNumber) // first input 1

	var secondNumber int64
	secondNumber = GetAndConvertToInt("Enter second number: ")
	fmt.Println(secondNumber) // second input 2

	thirdNumber := GetAndConvertToInt("Enter third number: ")
	fmt.Println(thirdNumber) // third input 3

	var fourthNumber int64
	fourthNumber = GetAndConvertToInt("Enter fourth number: ")
	fmt.Println(fourthNumber) //fourth input 4

	result := SumNumbers(firstNumber, secondNumber, thirdNumber, fourthNumber)
	fmt.Println("Result sum:" + strconv.FormatInt(result, 10)) // expected result 10
}

func GetAndConvertToInt(numberTitle string) int64 {
	var strNumber string
	fmt.Print(numberTitle)
	fmt.Scan(&strNumber)
	fmt.Println()
	numberResultI64, err := strconv.ParseInt(strNumber, 10, 64)
	if err != nil {
		fmt.Println("Your entered number is invalid")
		return GetAndConvertToInt(numberTitle)
	}
	return numberResultI64
}

func SumNumbers(a int64, b int64, c int64, d int64) int64 {
	var sum = a + b + c + d
	return sum
}
