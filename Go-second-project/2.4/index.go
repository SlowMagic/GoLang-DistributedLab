package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstPipe float64
	firstPipe = GetAndConvertToFloat("Enter first pipe:")
	fmt.Println(firstPipe) //5

	var secondPipe float64
	secondPipe = GetAndConvertToFloat("Enter second pipe:")
	fmt.Println(secondPipe) //6

	resultSide := PipesFilledByTime(firstPipe, secondPipe)
	fmt.Println("Pipes will fill for " + strconv.FormatFloat(resultSide, 'g', 3, 64))
}

func GetAndConvertToFloat(numberTitle string) float64 {
	var strNumber string
	fmt.Print(numberTitle)
	fmt.Scan(&strNumber)
	fmt.Println()
	numberResultI64, err := strconv.ParseFloat(strNumber, 64)
	if err != nil {
		fmt.Println("Invalid Number")
		return GetAndConvertToFloat(numberTitle)
	}
	return numberResultI64
}

func PipesFilledByTime(a float64, b float64) float64 {
	var result1 = 1/a + 1/b
	var result = 1 / result1

	return result // 2.727
}
