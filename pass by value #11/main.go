package main

import "fmt"

//pass by value
func changeName(x string) string {
	x = "ronoa"

	return x
}

//map memory is not duplicated when passed into function
func changeMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	name := "luffy"

	changeName(name)

	fmt.Println(name)

	menu := map[string]float64{
		"pie":           5.78,
		"coffee":        1.90,
		"hot chocolate": 2.78,
		"latte":         3.33,
	}

	changeMenu(menu)
	fmt.Println(menu)
}
