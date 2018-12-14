package main

import "fmt"

var z = 18

// declares there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
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
