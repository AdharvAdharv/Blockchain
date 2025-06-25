package main

import "fmt"

type Student struct{
	Name string
	Mark int
}

func getTopStudent (students []Student) []Student {
	var topStudents []Student
	for _,value :=range students{
		if value.Mark >80 {
			topStudents = append( topStudents,value)
		}
	}
     return topStudents
}

func main() {
	
	students := [] Student{
		{"Hiran",90},
		{"Varun",45},
		{"Bob",81},
	}
	toppers := getTopStudent(students)

	fmt.Println("Students with mark >80")
	fmt.Println(toppers)
}