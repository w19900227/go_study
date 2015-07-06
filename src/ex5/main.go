package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path : ", r.URL.Path)
	fmt.Println("scheme : ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key : ", k)
		fmt.Println("val : ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
	// fmt.Fprintf(w, "<h1>Hello %s!</h1><br><br><image src=/img/A.jpg />", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}
/*
package main

import (
    "fmt"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprintf(w, "<h1>Hello %s!</h1><br><br><image src=/img/A.jpg />", r.URL.Path[1:])
}

func img(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:]) 
}

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/img/", img)
    http.ListenAndServe(":80", nil)
}
*/
