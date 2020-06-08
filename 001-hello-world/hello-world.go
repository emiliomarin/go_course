package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	foo()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// Printing
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("What's up?")
}
