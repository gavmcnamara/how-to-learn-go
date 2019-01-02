package main

import "fmt"

//Create your own type. Have the
//underlying type be an int.
//create a VARIABLE of your new TYPE
//with the IDENTIFIER “x” using the “VAR” keyword
//in func main
//print out the value of the variable “x”
//print out the type of the variable “x”
//assign 42 to the VARIABLE “x” using the
// “=” OPERATOR
//print out the value of the variable “x”

type erin int

var gavin erin

func main() {

	fmt.Println(gavin)
	fmt.Printf("%T\n", gavin)
	gavin = 25
	fmt.Println(gavin)


}
