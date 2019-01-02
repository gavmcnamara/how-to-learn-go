package main

import "fmt"

func main() {

	fmt.Println("Start program")

	foo()

	fmt.Println("After foo")

	bar()

	fmt.Println("After bar")

	gavin()

	fmt.Println("End program!")
}

func foo() {

	fmt.Println("This is foo!")
}

func bar() {

	fmt.Println("This is bar!")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func gavin() {

	fmt.Println("This is gavin!")

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

// Control Flow
// (1) Sequence
// (2) loop; iterative
// (3) Conditional
