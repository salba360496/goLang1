//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func checkInt(intVal int) string {
	if (intVal%3 == 0) && (intVal%5 == 0) {
		return "FizzBuzz"
	} else if intVal%3 == 0 {
		return "Fizz"
	} else if intVal%5 == 0 {
		return "Buzz"
	}
	return ""
}

func main() {
	for i := 1; i <= 50; i++ {
		fmt.Println("The integer is", i, checkInt(i))
	}
}
