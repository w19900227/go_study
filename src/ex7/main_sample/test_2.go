package main_sample

import (    
        "fmt"
        "encoding/json"
        )
type data2 struct {
	Age int `json:"age"`
	Name *string `json:"name"`//加*用於判斷json mapping是否空
}
func Test2()  {
	body := []byte(`[
				{"age":6, "school":"school1", "name":"test_name"},
				{"age":7, "school":"school2"}
				]`)

    var data []data2

	json.Unmarshal(body, &data)

    for _, d := range data {
        if d.Name != nil {
            fmt.Printf("got age, %d : %s\n", d.Age, *d.Name)
        } else {
            fmt.Printf("got age, %d : \n", d.Age)
        }
    }
}

