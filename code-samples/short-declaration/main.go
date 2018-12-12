package main

import "fmt"

func main() {

	x := 42
	// := is the short declaration operator
	// this declares variable x and assigns a value
	// first time using a variable you have to declare it
	// x is declare and 42 is assign
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	// since x is already declared, just use '='
	y := 100 + 24
	// a statement is made up of expressions
	// 42 or 100 + 24 is an expression
	// y := 100 + 24 is a statement
	// '+' is an operator and '100', '24' are operands
	fmt.Println(y)
	// this program is made up of six statements
	z := "Bond, James"
	fmt.Println(z)

}
