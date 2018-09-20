package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Marice"] = "marice@gmail.com"
	emails["Diego"] = "diego@gmail.com"

	// Declare map and add kv
	// emails := map[string]string{"Bob": "bob@gmail.com", "Marice": "marice@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails)
}
