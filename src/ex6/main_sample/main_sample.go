package main_sample

type Sample_struct struct {
	Sample_1 string
}

func (this Sample_struct) MyFunc() string {
	return this.Sample_1
}


func MyFunc2() string {
	return "MyFunc2"
}
