package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement is", sum)
	}

	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("Continue is", sum)
		if sum%2 == 0 {
			continue
		}
	}

	for i := 0; i < 10; i++ {
		sum -= i
		fmt.Println("break is", sum)
		if sum<20 {
			break
		}
	}

}
