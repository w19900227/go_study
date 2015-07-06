package main

import (    
        "fmt"
         sam "./main_sample"
        )

func main() {
	x := sam.Sample_struct{Sample_1:"MyFunc"}
	myFunc := x.MyFunc()
	fmt.Println(myFunc)

	myFunc2 := sam.MyFunc2()
	fmt.Println(myFunc2)
}
