package main

import "./request"
import "fmt"

func main() {
	type D struct {
		Num int `json:"num"`
	}
	var d D
	d.Num = 82
	var request request.Request
	fmt.Println("--------GET--------")
	request.Get("http://localhost:9090/ex1")
	fmt.Println("--------POST--------")
	request.Post("http://localhost:9090/ex2", d)
	request.Post2("http://localhost:9090/ex2", d)
	fmt.Println("--------PUT--------")
	request.Put("http://localhost:9090/ex2", d)
	fmt.Println("--------DELETE-----")
	request.Delete("http://localhost:9090/ex2", d)
	
}

