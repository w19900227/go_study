package main

import "fmt"

func main() {
	a1, b1 := 1, 2
	func1(a1, b1)
	fmt.Println("a1 => ", a1)
	fmt.Println("-----------------------")

	a2 := []int{1, 2, 3}
	func2(a2)
	fmt.Println("a2 => ", a2)
	fmt.Println("-----------------------")

	a3 := 1
	func3(&a3)
	fmt.Println("a3 => ", a3)
	fmt.Println("-----------------------")
}

func func1(a1 ...int) {
	a1[0] = 3
	fmt.Println("func1 => ", a1)
}

func func2(a2 []int) {
	a2[0] = 3
	fmt.Println("func2 => ", a2)
}

func func3(a3 *int) {
	*a3 = 3
	fmt.Println("func3 => ", *a3)
}