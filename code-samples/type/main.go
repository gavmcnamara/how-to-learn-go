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

var b string
var c int

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

	zeroValue()
}

// primitive data types: basic or built-in data types
// bool, numberic, string, etc

// composite/aggregate data type: compose together data
// slice: hold many values of one type
// struct: hold many values of different types

func zeroValue() {
	fmt.Println("printing value of b:", b, "ending")
	fmt.Printf("%T\n", b)

	fmt.Println("printing value of c:", c, "ending")
	fmt.Printf("%T\n", c)

	b = "Bond, James Bond"
	c = 8

	fmt.Println("printing new value of b:", b)
	fmt.Printf("%T\n", b)

	fmt.Println("printing new value of c:", c)
	fmt.Printf("%T\n", c)
}

// Understanding zero value
// false for booleans
// 0 for integers
// 0.0 for floats
// "" for strings
// nil for: pointers, functions, interfaces,
// slices, channels, maps
