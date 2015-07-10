package main

import "fmt"


type Sample1 struct {}
type s1_interface interface {
  Coding1()
}
func (s1 *Sample1) Coding1() {
  fmt.Println("Coding 1")
}
func work1(s1 s1_interface) {
  s1.Coding1()
}

type Sample2 struct {}
type s2_interface interface {
  Coding2()
}
func (s2 Sample2) Coding2() {
  fmt.Println("Coding 2")
}
func work2(s2 s2_interface) {
  s2.Coding2()
}

func main() {

	// s1 := Sample1{}
	var s1 Sample1
	work1(&s1)

	// s2 := Sample2{}
	var s2 Sample2
	work2(s2)
}