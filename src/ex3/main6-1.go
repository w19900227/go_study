package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is : 12 * 2 = ", r1.area())
	fmt.Println("Area of r2 is : 9  * 4 = ", r2.area())
	fmt.Println("Area of c1 is : 10 * 10 * 3.14 = ", c1.area())
	fmt.Println("Area of c2 is : 25 * 25 * 3.14 = ", c2.area())

}

