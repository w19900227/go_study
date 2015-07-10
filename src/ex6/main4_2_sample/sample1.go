package main4_2_sample

import "fmt"

type Sample1 struct {}
type s1_interface interface {
  Coding1()
}
func (s1 *Sample1) Coding1() {
  fmt.Println("Coding 1")
}
func Work1(s1 s1_interface) {
  s1.Coding1()
}
