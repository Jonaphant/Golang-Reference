package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	//Declare and assign
	// fruitArr := [2]string{"Apple pie", "0range"}

	fruitSlice := []string{"Apple", "0range", "Raspberry", "Cantalope"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(fruitSlice)

	// Common built in array functions
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
