package main

import "fmt"

func main() {
	//strings
	var firstName string = "Mabo"
	fmt.Println(firstName)

	var secondName string = "John"
	fmt.Println(secondName)

	//var thirdName string

	fmt.Println(firstName, secondName)

	fourthName := "Mabo" //Short-hand assignment (cannot used outside a function)

	fmt.Println(fourthName)

	//integers
	var firstNum int = 12
	var secondNum int = 5
	thirdNum := 4

	fmt.Println(firstNum, secondNum, thirdNum)

	//bits & memory
	var firstNum1 int8 = 4 //signed 8 bi integer (between -128 to 127)
	fmt.Println(firstNum1)

	var secondNum1 int16 = 31000 //signed 16 bit integer (between -32768 to 32767)
	fmt.Println(secondNum1)

	var thirdNum1 uint = 25 //uint means the value cannot be less than 0
	fmt.Println(thirdNum1)

	//floating points
	var firstFloat float32 = 23.56634
	fmt.Println(firstFloat)
	var secondFloat float64 = 2378432623742389.56 //higher precision
	fmt.Println(secondFloat)

}
