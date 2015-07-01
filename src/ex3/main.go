package main

import "fmt"

type person struct {
	name string
	age int
}

func personeFunc(p person) (string, int) {
	return p.name, p.age
}

func main() {
	var tom person
	tom.name, tom.age = "tom", 18

	bob := person{age:25, name:"bob"}

	paul := person{"paul", 34}

	tom_name, tom_age := personeFunc(tom)
	bob_name, bob_age := personeFunc(bob)
	paul_name, paul_age := personeFunc(paul)

	fmt.Printf("tom - name : %s, age : %d\n", tom_name, tom_age)
	fmt.Printf("bob - name : %s, age : %d\n", bob_name, bob_age)
	fmt.Printf("paul - name : %s, age : %d\n", paul_name, paul_age)

}

