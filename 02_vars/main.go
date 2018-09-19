package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// ------ Using variables ------
	// var name string = "Jonathan"
	// var age int = 22

	//When values are inferred as one of the basic data types, they don't require the type specification
	var name = "Jonathan"
	var age int32 = 22

	//Short hand of creating variables. Can only be used inside a function
	name2 := "Jonathan2"
	size := 1.5

	name3, size2 := "Jonathan3", 1.6

	//Similar to JS, const variables cannot be reassigned values
	const isCool = true

	fmt.Println(name, age, name2, isCool, name3, size2)
	fmt.Printf("%T\n", size)

}
