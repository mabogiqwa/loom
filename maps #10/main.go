package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":   4.99,
		"pie":    7.99,
		"salad":  6.99,
		"coffee": 4.05,
	} //maps work similarly to dictionaries

	fmt.Println(menu)
	fmt.Println(menu["soup"]) //prints value of the key soup

	//looping maps
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	//ints as key type
	phonebook := map[int]string{
		23423245: "gee",
		26439653: "vee",
		23242345: "bee",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[23423245])
}
