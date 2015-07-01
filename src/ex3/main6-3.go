package main

import "fmt"

type Human struct {
	name, phone string
	age int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (s Student) sayHi() {
	fmt.Printf("Hi i am %s you can call me on %s\n", s.name, s.phone)
}

func (e Employee) sayHi() {
	fmt.Printf("Hi i am %s, i work at %s. call me on %s\n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"mark", "0911-222-333", 18}, "test school"}
	tom := Employee{Human{"tom",  "0944-555-666", 28}, "test company"}

	mark.sayHi()
	tom.sayHi()
}

