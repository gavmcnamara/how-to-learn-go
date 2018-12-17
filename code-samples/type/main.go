package main

import (
	"fmt"
)

var y = 42

// DECLARED the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "shaken, not stirred"

var z = "Shaken, not stirred"

var a string = `James said,
"Shaked, 

not stirred"`

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// NOT a DYNAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

// primitive data types: basic or built-in data types
// bool, numberic, string, etc

// composite/aggregate data type: compose together data
// slice: hold many values of one type
// struct: hold many values of different types
