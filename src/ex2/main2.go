package main

import "fmt"

//宣告一個函數類型
type testInt func(int) bool

func isOdd(var1 int) bool {
	if var1%2 == 0 {
		return false
	}
	return true
}

func isEvent(var1 int) bool {
	if var1%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd element of slice are = ", odd)
	event := filter(slice, isEvent)
	fmt.Println("Event element of slice are = ", event)

}

