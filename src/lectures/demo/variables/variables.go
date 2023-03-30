package main

import "fmt"

func main() {
	var myName = "Salman"
	fmt.Println("My name is ",myName)

	var lastName string  = "Ahmed"
	fmt.Println("My Last Name is",lastName)

	userName := "Salba"
	fmt.Println("My userName is",userName)

	var num int
	fmt.Println("My number is",num)

	age,dob := 37,"15/10/1986"
	fmt.Println("My age is",age,"and date of birth is",dob )

	part1,other := 5,5
	fmt.Println("part1",part1,"other",other)

	part2,other:= 2,3
	fmt.Println("part2",part2,"other",other)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("LessonName =",lessonName,"LessonType =",lessonType)

	sum:= part1+part2
	fmt.Println("The sum of two parts is",sum)

	word1,word2, _ := "Hello","World","!"

	fmt.Println(word1,word2)

}
