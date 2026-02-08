package main

import "fmt"

//non-pointer objects: strings, ints, floats, booleans, arrays, structs
//pointer-wrapper vals: slices, maps, functions

func changeValue(x string) {
	x = "x^3dxnc$3#"
}

func main() {
	name := "flint"

	changeValue(name)

	fmt.Println("Memory address of the variable name is: ", &name)
	m := &name //variables m and name share the same memory addresses
	fmt.Println(name)
	fmt.Println(*m)

	*m = "stone"
	fmt.Println(name)
}
