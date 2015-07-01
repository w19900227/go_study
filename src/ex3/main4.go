package main

import "fmt"

type person struct {
	name string
	age int
}

type Student struct {
	person
	school string
	name string
}

func main() {

	mark := Student{person{"tom", 18}, "ntpu", "Student name"}

	fmt.Println("His person name is : ", mark.person.name)
	fmt.Println("His person age is : ", mark.person.age)
	fmt.Println("His school is : ", mark.school)
	fmt.Println("His name is : ", mark.name)

}

