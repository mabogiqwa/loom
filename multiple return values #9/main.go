package main

import (
	"fmt"
	"strings"
)

func getInitials(fullName string) (string, string) {
	s := strings.ToUpper(fullName)
	names := strings.Split(s, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	firstName, secondName := getInitials("Mabo Giqwa")
	fmt.Println(firstName, secondName)

	firstName1, secondName1 := getInitials("Yuji Hatake")
	fmt.Println(firstName1, secondName1)

	firstName2, secondName2 := getInitials("Jane")
	fmt.Println(firstName2, secondName2)

}
