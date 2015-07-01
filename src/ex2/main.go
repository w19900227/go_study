package main

import "fmt"


func main() {
	var1 := 10
	var2 := 5

	sum, subtract := testFunc(var1, var2)
	fmt.Printf("%d + %d = %d\n", var1, var2, sum)
	fmt.Printf("%d - %d = %d\n", var1, var2, subtract)
}

func testFunc(inputVar1 int, inputVar2 int) (sum int, subtract int) {
	sum = inputVar1 + inputVar2
	subtract = inputVar1 - inputVar2
	return 
}
