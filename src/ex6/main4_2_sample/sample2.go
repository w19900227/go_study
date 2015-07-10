package main4_2_sample

import "fmt"

type Sample2 struct {}
type s2_interface interface {
  Coding2()
}
func (s2 Sample2) Coding2() {
  fmt.Println("Coding 2")
}
func Work2(s2 s2_interface) {
  s2.Coding2()
}
