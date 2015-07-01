package main

import "fmt"

type funcName func() string

func ex1() string {
	return "ex1"
}

func ex2() string {
	return "ex2"
}

func filter(f funcName) string {
	return f()
}

func main() {
	
	string_ex1 := filter(ex1)
	fmt.Printf("ex1 return => %s\n", string_ex1)

	string_ex2 := filter(ex2)
	fmt.Printf("ex2 return => %s\n", string_ex2)
}

