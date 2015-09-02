package request

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
)
import "fmt"

var types string = "application/json"

type Request struct {

}

func (this *Request) Get(url string) []byte {
    response, _ := http.Get(url)
    datas, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(datas))
    return datas
}
func (this *Request) Post(url string, data interface{}) {
    client := &http.Client{}
    
    buf, _ := json.Marshal(data)
    body := strings.NewReader(string(buf))

    response, _ := http.NewRequest("POST", url, body)
    response.Header.Set("Content-Type", "application/json")
    re, _ := client.Do(response)
    defer re.Body.Close()
    datas, _ := ioutil.ReadAll(re.Body)
    // fmt.Println(datas)
    fmt.Println(string(datas))
    
}

func (this *Request) Post2(url string, data interface{}) {
    client := &http.Client{}
    body := strings.NewReader(`{"num":22}`)

    response, _ := http.NewRequest("POST", url, body)
    response.Header.Set("Content-Type", "application/json")
    re, _ := client.Do(response)
    defer re.Body.Close()
    datas, _ := ioutil.ReadAll(re.Body)
    // fmt.Println(datas)
    fmt.Println(string(datas))
    
}

func (this *Request) Put(url string, data interface{}) {
    client := &http.Client{}
    body := strings.NewReader(`{"num":22}`)

    response, _ := http.NewRequest("PUT", url, body)
    response.Header.Set("Content-Type", "application/json")
    re, _ := client.Do(response)
    defer re.Body.Close()
    datas, _ := ioutil.ReadAll(re.Body)
    // fmt.Println(datas)
    fmt.Println(string(datas))
    
}


func (this *Request) Delete(url string, data interface{}) {
    client := &http.Client{}
    body := strings.NewReader(`{"num":22}`)

    response, _ := http.NewRequest("DELETE", url, body)
    response.Header.Set("Content-Type", "application/json")
    re, _ := client.Do(response)
    defer re.Body.Close()
    datas, _ := ioutil.ReadAll(re.Body)
    // fmt.Println(datas)
    fmt.Println(string(datas))
    
}