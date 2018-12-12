package main

// declaration of package
// this code belogns in package main

// Packages allows us to organize our code
// into differnt packages, similar to how
// folders are organized on our computers

import (
	"fmt"
)

// importing packages that you USE in your code
func main() {

	fmt.Println("Hello Everyone", 42, true)
	// from package fmt im calling Println

	n, err := fmt.Println("test", 13, false)
	fmt.Println(n)
	fmt.Println(err)
	// catch any returns
	// n = number of bytes written
	// err = any errors
	// NEVER HAVE A VARIABLE YOU DONT USE
	// if need to hold space as variable use _
	// ex) n, _
	// but will not be able to use fmt.Println(_)
}

// fmt.Println can pass in value of any type
// String
// Integer
// bool
