package main

import "fmt"

func main() {
	x := 5
	y := 10

	// If
	// No parenthesis around the condition are typically the standard.
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// Else if
	color := "blue"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Printf("Color is blue")
	} else {
		fmt.Printf("Color is neither blue or red")
	}

	fmt.Println()
	// Switch

	color = "red"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is neither blue or red")
	}
}
