package main

import (
	"fmt"
	sam2 "./main2_sample"
)

func main() {
	val := sam2.Sample2_struct{"Sample2_struct"}
	var m2 sam2.Main2_interface
	
	m2 = val

	m2.MyFunc()
	fmt.Println(m2.MyFunc2())

}
