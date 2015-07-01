package main

import "fmt"

func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func main() {
	s := gen(1, 2, 3, 5)

	for a := range s {
		fmt.Println(a)
	}

    c := make(chan int)
    go func() { 
        c <- 42 
        close(c)
    }()
    val := <-c
    println(val)
}

