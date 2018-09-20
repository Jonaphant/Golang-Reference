package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// User * to read val from address
	fmt.Printf("%d\n", *b)
	fmt.Printf("%d\n", *&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
