package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {

	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz1 is greater then quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz1 is less then quiz2")
	} else {
		fmt.Println("Quiz1 is equal to quiz2")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Scores Acceptable")
	} else {
		fmt.Println("Scores are not acceptable")
	}
}
