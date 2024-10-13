package main

import (
	"fmt"
	"time"
)

func main() {
	//task1
	dateNow := time.Now()
	fmt.Print("\ntask1\n", "Date/time today is: ", dateNow.Format(time.RFC850))

	myInteger := int(1234)
	myFloat := float64(12.34)
	myString := string("1234")
	myBool := bool(float64(myInteger) > myFloat)

	fmt.Print(
		"\n\ntask2\n", "int: ", myInteger, "\n",
		"float64: ", myFloat, "\n",
		"string: ", myString, "\n",
		"bool: ", myBool, "\n")

	IntegerOperation(myInteger, int(myFloat))
	FloatOperation(float64(myInteger), myFloat)
	AverageValue(myInteger, int(myFloat), myFloat)
}

func IntegerOperation(a int, b int) {
	fmt.Print(
		"\ntask4\n", "summ integer: ", a+b, "\n",
		"multiply integer: ", a*b, "\n",
		"dividing integer: ", a/b, "\n",
		"substraction of integers: ", a-b, "\n")
}

func FloatOperation(a float64, b float64) {
	fmt.Print(
		"\ntask5\n", "summ float64: ", a+b, "\n",
		"substraction of float64: ", a-b, "\n")
}

func AverageValue(a int, b int, c float64) {
	average := (float64(a) + float64(b) + c) / 3
	fmt.Println("\ntask6\n", "Average value: ", average)
}
