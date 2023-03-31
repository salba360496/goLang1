package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap price")
	case p < 10:
		fmt.Println("Moderate price")
	default:
		fmt.Println("Very expensiive price")
	}
	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("FirstClass seating")
	default:
		fmt.Println("Other seating")
	}
}
