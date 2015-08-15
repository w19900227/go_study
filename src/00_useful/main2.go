package main

import "fmt"
import "runtime"
// import "encoding/json"


func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}
