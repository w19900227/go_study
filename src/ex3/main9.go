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

func (s Student) test1() {
	fmt.Println("myfunc to test1")
}

func (e Employee) test2() {
	fmt.Println("myfunc to test2")
}

type Men interface {
	sayHi()
	test1()
}

type Young interface {
	sayHi()
	test2()
}

func main() {
	mark := Student{Human{"mark", "0911-222-333", 18}, "test school"}
	tom := Employee{Human{"tom",  "0944-555-666", 28}, "test company"}

	var i Men
	i = mark
	i.sayHi()
	i.test1()

	var y Young
	y = tom
	y.sayHi()
	y.test2()
}

