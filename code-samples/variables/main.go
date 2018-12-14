package main

import "fmt"

var z = 18

func main() {
	// short declaration operator
	// DECLARE variable and ASSIGN a value
	x := 42
	fmt.Println(x)

	// Whenever needed to create variable outside function
	// body use var and NOT short declaration operator
	var y = 345
	fmt.Println(y)

	fmt.Println(z)
}

// NOTE: limit scope of your varibles and try to use
// short declaration operator as much as possible!