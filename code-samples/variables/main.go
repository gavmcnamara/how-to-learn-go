package main

import "fmt"

// DECLARE the variable "z"
// ASSIGN  the value 43
// declare & assign = initialization
var z = 18

// declares there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for intigers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var y int

func main() {
	// short declaration operator
	// DECLARE variable and ASSIGN a value
	x := 42
	fmt.Println(x)

	// Whenever needed to create variable outside function
	// body use var and NOT short declaration operator
	fmt.Println(z)

	foo()

	fmt.Println(y)
}

func foo() {

	fmt.Println("again:", z)
}

// NOTE: limit scope of your varibles and try to use
// short declaration operator as much as possible!
