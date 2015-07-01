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

func (h Human) sayHi() {
	fmt.Printf("Hi i am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	mark := Student{Human{"mark", "0911-222-333", 18}, "test school"}
	tom := Employee{Human{"tom",  "0944-555-666", 28}, "test company"}

	mark.sayHi()
	tom.sayHi()
}

