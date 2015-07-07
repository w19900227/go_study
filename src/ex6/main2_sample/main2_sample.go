package main2_sample

import "fmt"

type Main2_interface interface {
	MyFunc()
	MyFunc2() string
}

type Sample2_struct struct {
	Sample_2 string
}

func (this Sample2_struct) MyFunc() {
	// return this.Sample_2
	fmt.Println(this.Sample_2)
}


func (this Sample2_struct) MyFunc2() string {
	return "MyFunc2"
}
