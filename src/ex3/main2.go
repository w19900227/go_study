package main

import "fmt"

type person struct {
	name string
	age int
}

type Student struct {
	person
	school string
}

func main() {

	mark := Student{person{"tom", 18}, "ntpu"}

	fmt.Println("His name is : ", mark.name)
	fmt.Println("His age is : ", mark.age)
	fmt.Println("His school is : ", mark.school)

	fmt.Println("=====================")

	mark.age += 7
	mark.school = "ntu"

	fmt.Println("His name is : ", mark.name)
	fmt.Println("His age is : ", mark.age)
	fmt.Println("His school is : ", mark.school)

}

