package main

import (
	"fmt"
)

var y = 42

// DECLARED the VARIABLE with the IDENTIFIER "z"
// is of TYPE string 
// and I ASSIGN the VALUE "shaken, not stirred"

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// NOT a DYNAMIC programming language

var z = "Shaken, not stirred"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
