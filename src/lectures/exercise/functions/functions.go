//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greeting(name string) {
	fmt.Println("Hello", name, "welcome to the tutorial of golang!")
}

// func fullName(firstName ,lastName string)(string,string){
// 	return firstName,lastName
// }

func firstName(firstName string) string {
	return firstName
}

func add(a, b, c int) int {
	return a + b + c
}

func oneNumberReturn() int {
	return 2
}

func twoNumberReturn() (int, int) {
	return 2, 3
}

func main() {
	greeting("Salman")
	fmt.Println("Hi there", firstName("Salman"))
	fmt.Println("Addition of three numbers", add(2, 3, 4))
	x, y := twoNumberReturn()
	fmt.Println(add(oneNumberReturn(), x, y))
}
