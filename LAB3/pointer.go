package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Marks float64
}

func modifyVariable(num *int) {
	var temp int
	fmt.Print("Enter a new value for a: ")
	fmt.Scanln(&temp)
	*num = temp
}
func addDataintoStructure(s *Student) {
	fmt.Print("Enter Name: ")
	fmt.Scanln(&s.Name)
	fmt.Print("Enter Age: ")
	fmt.Scanln(&s.Age)
	fmt.Print("Enter Marks: ")
	fmt.Scanln(&s.Marks)
}
func main() {
	a := 10

	fmt.Println("Address of a is: ", &a)
	p := &a
	fmt.Println(*p)
	fmt.Println("To modify the value of a, we can use the pointer p")
	var input int
	fmt.Print("Enter a new value for a: ")
	fmt.Scanln(&input)

	modifyVariable(&input)
	fmt.Println("Value of a after modification: ", input)
	fmt.Println("To make the student struct ")
	s1 := new(Student)
	addDataintoStructure(s1)
	fmt.Println(*s1)

}
