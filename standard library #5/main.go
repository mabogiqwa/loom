package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var sen string = "Hello World"

	sen2 := "Hello Earthlings"

	fmt.Println(sen)
	fmt.Println(sen2)

	fmt.Println(strings.Contains(sen, "Hello"))         //Checks if the substring "Hello" is in the variable sen
	fmt.Println(strings.ReplaceAll(sen, "Hello", "Hi")) //Replaces all instances of the substring "Hello" with "Hi"
	fmt.Println(strings.ToUpper(sen))                   //Transforms all characters to upper case
	fmt.Println(strings.Split(sen, " "))                //splits words in a string into an array
	ages := []int{45, 66, 12, 45, 8, 63, 12, 90, 71}

	sort.Ints(ages) //sorts array in ascending
	fmt.Println(ages)

	index := sort.SearchInts(ages, 8)
	fmt.Println(index)

	variousNames := []string{"Mabo", "Yamamoto", "Luigi", "Goku", "Yuji"}

	sort.Strings(variousNames) //Sorts according to alphabetical order
	fmt.Println(variousNames)

	//fmt.Println("Original string:", sen)
}
