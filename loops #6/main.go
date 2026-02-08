package main

import "fmt"

func main() {
	/*
		x := 0
		for x < 5 {
			fmt.Println("Value of x is:", x)
			x++
		}
	*/
	//Alternative definition
	for i := 0; i < 5; i++ {
		fmt.Print("Value of i is: ", i, "| ")
	}
	fmt.Println()

	names := []string{"Mabo", "Yamamoto", "Luigi", "Goku", "Yuji"}
	/*
		for i := 0; i < len(names); i++ {
			fmt.Println(names[i])
		}
	*/

	for index, value := range names {
		fmt.Printf("The position at index %v is %v\n", index, value)
	}
}
