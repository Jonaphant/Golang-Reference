package main

import "fmt"

func main() {
	ids := []int{34, 62, 66, 33, 83, 21, 5}

	// Loop thruogh the ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add all the ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum is %d\n", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Marice": "marice@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s's email is %s\n", k, v)
	}
}
