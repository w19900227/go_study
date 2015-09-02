package main

import (
	"net/http"
	"encoding/json"
	"log"
    // "io/ioutil"
)
import "fmt"
type getNum struct {
	Num int `json:"num"`
}
func ex1_func(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
    type head struct {
    	Nums []int `json:"nums"`
    }
    var tt []int
    for i:=1; i<=5; i++ {
        tt = append(tt, i)
    }
    aa := head{tt}

    js, _ := json.Marshal(aa)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ex2_func(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	
	var n getNum

	decoder := json.NewDecoder(r.Body)
    decoder.Decode(&n)
	
    type head struct {
    	Nums []int `json:"nums"`
    }

    var tt []int
    for i:=1; i<=n.Num; i++ {
        tt = append(tt, i)
    }
    aa := head{tt}

    js, _ := json.Marshal(aa)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", ex1_func)
	http.HandleFunc("/ex1", ex1_func)
    http.HandleFunc("/ex2", ex2_func)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}

