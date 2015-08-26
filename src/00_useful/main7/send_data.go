package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

func main() {
	fmt.Println("ex1 ---------------------------")
	ex1()
	fmt.Println("ex2 ---------------------------")
	ex2()
}

func ex1() {
	url := "http://localhost:9090/ex1"

	resp, _ := http.Get(url)	
	body, _ := ioutil.ReadAll(resp.Body)

	type Nums struct {
		Num []int `json:"nums"`
	}

	var data Nums
	json.Unmarshal(body, &data)

	fmt.Println(data)
	fmt.Println(data.Num[0])
}

func ex2() {
	type Person struct {
		Num int `json:num`
	}

	url := "http://localhost:9090/ex2"
	types := "application/json"
	n := 50

	person := &Person{n}
	buf, _ := json.Marshal(person)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post(url, types, body)
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))
}