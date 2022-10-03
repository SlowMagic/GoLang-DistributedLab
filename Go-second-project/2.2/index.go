package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a float64
	a = GetAndConvertToFloat("Enter side (A) of a triangle:")
	fmt.Println(a) //15

	var b float64
	b = GetAndConvertToFloat("Enter side (B) of a triangle:")
	fmt.Println(b) //12

	resultHypotenuse := TriangleHypotenuseCalc(a, b)
	fmt.Println("Hypotenuse is " + strconv.FormatFloat(resultHypotenuse, 'g', 10, 64)) // 19.20937271

	resultSquare := TriangleSquareCalc(a, b)
	fmt.Println("Square is " + strconv.FormatFloat(resultSquare, 'g', 10, 64)) // 90

}

func GetAndConvertToFloat(numberTitle string) float64 {
	var strNumber string
	fmt.Print(numberTitle)
	fmt.Scan(&strNumber)
	fmt.Println()
	numberResultFloat, err := strconv.ParseFloat(strNumber, 64)
	if err != nil {
		fmt.Println("Invalid Number")
		return GetAndConvertToFloat(numberTitle)
	}
	return numberResultFloat
}
func TriangleHypotenuseCalc(a float64, b float64) float64 {
	var result = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return result
}

func TriangleSquareCalc(a float64, b float64) float64 {
	var result = 0.5 * a * b
	return result
}
