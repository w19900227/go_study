package main

import sam4_2 "./main4_2_sample"

func main() {

	// s1 := Sample1{}
	var s1 sam4_2.Sample1
	sam4_2.Work1(&s1)

	// s2 := Sample2{}
	var s2 sam4_2.Sample2
	sam4_2.Work2(s2)
}