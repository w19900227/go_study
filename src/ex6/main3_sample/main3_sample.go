package main3_sample

import "fmt"

type Main3_interface interface {
	MyFunc()
	MyFunc2() string
}

type Sample3_struct struct {
	Sample_3 string
}

func (this Sample3_struct) MyFunc() {
	// return this.Sample_3
	fmt.Println(this.Sample_3)
}


func (this Sample3_struct) MyFunc2() string {
	return "MyFunc2"
}


type Main3_interface_2 interface {
	MyFunc()
	MyFunc2() string
}

type Sample3_struct_2 struct {
	Sample_3 string
}

func (this Sample3_struct_2) MyFunc() {
	// return this.Sample_3
	fmt.Println(this.Sample_3)
}


func (this Sample3_struct_2) MyFunc2() string {
	return "MyFunc2"
}
