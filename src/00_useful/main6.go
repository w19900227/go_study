package main

import (
	"net/http"
	"encoding/json"
	"log"
	// "fmt"
)

type ex1 ex1_1
type ex1_1 struct {
	E1 []string
}

func ex1_func(w http.ResponseWriter, r *http.Request) {
	e1 := ex1{[]string{"ex1-1","ex1-2"}}

	js, err := json.Marshal(e1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ex5_func(w http.ResponseWriter, r *http.Request) {
    type List struct {
        Service []string `json:"service"`
    }

    var list List
    list.Service = append(list.Service, "a")
    list.Service = append(list.Service, "b")

    js, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type ex2 []ex2_1
type ex2_1 struct {
	E2 string
}

func ex2_func(w http.ResponseWriter, r *http.Request) {
	e2 := ex2{{"ex2-1"}, {"ex2-2"}}

	js, err := json.Marshal(e2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}


func ex3_func(w http.ResponseWriter, r *http.Request) {
    var Datas []interface{}
    type Data struct {
        Name string `json:"name"`
        Age int `json:"age"`
    }

    for i:=1; i<3; i++ {
        Datas = append(Datas, Data{"ex3", i})
    }

    ret, err := json.Marshal(Datas)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(ret)
}

func ex4_func(w http.ResponseWriter, r *http.Request) {
    type Data struct {
        Name string `json:"name"`
        Age int `json:"age"`
    }
    type Datas []Data 
        
    result := Datas{}
    var data Data
    for i:=1; i<3; i++ {
        data.Name = "ex4"
        data.Age = i
        result = append(result, data)
        
    }

    ret, err := json.Marshal(result)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(ret)
}

func main() {
	http.HandleFunc("/", ex4_func)
	http.HandleFunc("/ex1", ex1_func)
	http.HandleFunc("/ex2", ex2_func)
	http.HandleFunc("/ex3", ex3_func)
	http.HandleFunc("/ex4", ex4_func)
	http.HandleFunc("/ex5", ex5_func)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}

