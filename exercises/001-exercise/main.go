package main

import "fmt"

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
