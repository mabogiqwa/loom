package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 20 {
		fmt.Println("Age is strictly less than 20")
	} else if age < 30 {
		fmt.Println("Age is strictly less than 30")
	} else if age < 40 {
		fmt.Println("Age is strictly less than 40")
	} else {
		fmt.Println("Age does not satisfy any of the requirements")
	}

	names := []string{"mario", "mano", "yoshi", "pear", "orange"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue //skips onto next iteration
		}
		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break //exits loop
		}

		fmt.Printf("The value at pos %v is %v\n", index, value)
	}
}
