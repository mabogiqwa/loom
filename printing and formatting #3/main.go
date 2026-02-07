package main

import "fmt"

func main() {
	age := 65
	name := "Shaun"

	fmt.Print("Hello, ")
	fmt.Print("World \n")
	fmt.Print("Onto a new line.")

	fmt.Println()

	fmt.Println("My name is,", name, ", and my age is", age)

	//Printf
	fmt.Printf("My name is %v and my age is %v\n", name, age)

	name1 := "Samantha"
	age1 := 34

	fmt.Printf("My name is %q and my age is %q\n", name1, age1)

	fmt.Printf("The age is of type %T \n", age)

	fmt.Printf("The score is %f points\n", 204.67)

	fmt.Printf("The score was %0.1f points \n", 255.756383)

	//Saving string to a variable
	var str = fmt.Sprintf("My name is %v and my age is %v\n", name, age)
	fmt.Println("The string stored in str is:", str)
}
