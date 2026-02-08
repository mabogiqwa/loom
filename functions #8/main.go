package main

import (
	"fmt"
	"math"
)

func add(num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

func substract(num1 int, num2 int) {
	fmt.Println(num1 - num2)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	add(4, 5)
	substract(9, 8)
	fmt.Printf("Circle area is %0.3f", circleArea(567.9765))
}
