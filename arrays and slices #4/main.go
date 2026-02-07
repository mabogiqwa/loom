package main

import "fmt"

func main() {
	//var weights [5]float64 = [5]float64{65.78, 55.65, 81.432, 78.90, 121.34} //array length cannot be changed

	var weights = [5]float64{65.78, 55.65, 81.432, 78.90, 121.34} //different way to instantiate array

	fmt.Println(weights, len(weights))

	var names = [3]string{"Donald", "New York", "Espana"}
	fmt.Println(names, len(names))
	names[0] = "Isaac"
	fmt.Println(names, len(names))

	//slices - size is changeable
	var scores = []int{90, 10, 30}
	scores[1] = 66
	scores = append(scores, 109)
	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[0:3] //from index 0 to index 2
	fmt.Println(rangeOne)

	rangeTwo := names[1:] //From index 1 onwards
	fmt.Println(rangeTwo)

	rangeThree := names[:3] //From index 0 till index 2
	fmt.Println(rangeThree)

	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
