package main

import "fmt"

func main() {

	sum := adder()

	for i := 0; i < 10; i++ {
		sum(5)
	}
}

func adder() func(int) {
	sum := 0
	return func(x int) {
		sum += x
		fmt.Printf("Sum is now %d\n", sum)
	}
}
