package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type ex1 ex1_1
type ex1_1 struct {
	Test []string
}

func ex1_func(w http.ResponseWriter, r *http.Request) {
	e1 := ex1{[]string{"aaa","bbb"}}

	js, err := json.Marshal(e1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type ex2 []ex2_1
type ex2_1 struct {
	Name string
}

func ex2_func(w http.ResponseWriter, r *http.Request) {
	e2 := ex2{{"aa"}, {"bb"}}

	js, err := json.Marshal(e2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}


func ex3_func(w http.ResponseWriter, r *http.Request) {
    var ta []interface{}
    type tt struct {
        Name string `json:"name"`
        Age int `json:"age"`
    }
    for i:=1; i<3; i++ {
        ta = append(ta, tt{"aaa", i})
        // tt.Name = "aaa"
        // tt.Age = i
        // ta = append(ta, tt)
    }
    // ta := tt{"aaa", 2}
    // fmt.Println(ta)

    js, err := json.Marshal(ta)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func main() {
	http.HandleFunc("/", ex1_func)
	http.HandleFunc("/ex2", ex2_func)
	http.HandleFunc("/ex3", ex3_func)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}

