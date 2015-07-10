package main

import (
	"fmt"
	sam3 "./main3_sample"
)

func test(inter *interface) interface {
	return sam3.inter
}

func main() {
	val1 := sam3.Sample3_struct{"Sample3_struct"}
	val2 := sam3.Sample3_struct_2{"Sample3_struct_2"}
	// var m3 sam3.Main_interface
	var m3 test(Main_interface)
	m3 = val1

	m3.MyFunc()
	fmt.Println(m3.MyFunc2())

}
