package main

import "fmt"

//Using the short declaration operator,
//ASSIGN these VALUES to VARIABLES with
//the IDENTIFIERS “x” and “y” and “z”
//42
//“James Bond”
//true
//Now print the values stored in those variables using
//a single print statement
//ultiple print statements

func main() {

	x, y, z := 42, "James Bond", true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T", z)
}
