package main

import (    
        "fmt"
        "encoding/json"
        "./main_sample"
        )
type test struct {
	Name string `json:"name"`
}

func test1() {

	body := []byte(`{"name":"Wednesday", "Age":6, "Parents":["Gomez","Morticia"]}`)

    var data test
	json.Unmarshal(body, &data)
	fmt.Println(data.Name)
}

func main() {
	test1()
	main_sample.Test2()
}


