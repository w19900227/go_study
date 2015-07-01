package main

import "fmt"

type Skills []string

type person struct {
	name string
	age int
}

type Student struct {
	person
	Skills
	school string
}

func main() {

	mark := Student{person:person{"tom", 18}, school:"ntpu"}

	fmt.Println("His name is : ", mark.name)
	fmt.Println("His age is : ", mark.age)
	fmt.Println("His school is : ", mark.school)

	mark.Skills = []string{"anatomy"}
	fmt.Println("His skills are : ", mark.Skills)
	
	mark.Skills = append(mark.Skills, "play", "music")
	fmt.Println("His skills are : ", mark.Skills)

}

