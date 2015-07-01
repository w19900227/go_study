package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name, phone string
	age int
}

func (h Human) String() string {
	return " " + h.name + " - " + strconv.Itoa(h.age) + " years - " + h.phone 
}

type Stringer interface {
	String() string
}

func main() {
	mark := Human{"mark", "0911-222-333", 18}
	fmt.Println("This Human is : ", mark)
}

