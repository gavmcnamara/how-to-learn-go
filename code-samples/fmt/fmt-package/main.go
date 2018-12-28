package main

import "fmt"

var y = 42

// %T = type, %b = binary, %#x = hexidecimal
// %x = regular hexidecimal (base 16, with lower case a-f)
func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	// fmt.Sprintf prints to string
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)

	// #v the value in a default format
	fmt.Printf("%v", y)
}
