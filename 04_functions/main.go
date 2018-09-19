package main

import "fmt"

//Function format:
//func <function name>(<parameter name> <parameter type>) <return type> {}

func greeting(name string) string {
	return "Hello " + name + "!"
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Jonathan"))
	fmt.Println(getSum(5, 6))
}
