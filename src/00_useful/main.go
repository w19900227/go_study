package main

import "fmt"
import "encoding/json"


func main() {
	fmt.Println("ex1 ---------------------------")
	ex1()
	fmt.Println("ex2 ---------------------------")
	ex2()

}

func ex1() {
	var f interface{}
	json_data := []byte(`{"Name":"Wednesday", "Age":6, "Parents":["Gomez","Morticia"]}`)

	json.Unmarshal(json_data, &f)
	m := f.(map[string]interface{})
	for k, v := range m {
	    switch vv := v.(type) {
	    case string:
	        fmt.Println("is string -> ", vv)
	    case int:
	        fmt.Println("is int -> ", vv)
	    case []interface{}:
	        fmt.Print("is interface -> ")
	        fmt.Printf("%+v\n", v)
	    default:
	        fmt.Println("is of a type I don't know how to handle -> ", k)
	    }
	}
}
func ex2() {

	var f interface{}

	f = map[string]interface{}{
	    "Name": "Wednesday",
	    "Age":  6,
	    "Parents": []interface{}{
	        "Gomez",
	        "Morticia",
	    },
	}
	m := f.(map[string]interface{})

	for k, v := range m {
	    switch vv := v.(type) {
	    case string:
	        fmt.Println("is string -> ", vv)
	    case int:
	        fmt.Println("is int -> ", vv)
	    case []interface{}:
	        fmt.Print("is interface -> ")
	        fmt.Printf("%+v\n", v)
	    default:
	        fmt.Println("is of a type I don't know how to handle -> ", k)
	    }
	}
}

func ex3() {
	var jsonBlob = []byte(`[
        {"Name": "Platypus"},
        {"Name": "Quoll",    "Order": 100}
    ]`)
    type Animal struct {
        Name  string 
        Order *int   
    }
    var animals []Animal
    err := json.Unmarshal(jsonBlob, &animals)
    if err != nil {
        fmt.Println("error:", err)
    }
    for _, a := range animals {
        if a.Order != nil {
            fmt.Printf("got order, %s : %d\n", a.Name, *a.Order)
        }
    }
}