package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	//Alternative way
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// HasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Jon", lastName: "Doe", city: "Boston", gender: "m", age: 25}

	// Alternative
	person2 := Person{"Samantha", "Smith", "Boston", "f", 50}

	fmt.Println(person1)

	// Using the pointer reciever method
	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Lee")

	//Using the value reciever method
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
