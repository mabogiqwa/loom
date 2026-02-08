package main

import (
	"fmt"
	"strings"
)

func main() {
	var sen string = "Hello World"

	sen2 := "Hello Earthlings"

	fmt.Println(sen)
	fmt.Println(sen2)

	fmt.Println(strings.Contains(sen, "hello"))
}
