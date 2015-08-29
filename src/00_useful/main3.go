package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"
// import "container/list"


type promotion_block struct {
	Status string `json:"status"`
	Data []block_array `json:"data"`
}

type block_array struct {
	Id string `json:"id"` 
	Country_code string `json:"country_code"` 
	Title string `json:"title"` 
	Duration_from string `json:"duration_from"` 
	Duration_to string `json:"duration_to"` 
	Types string `json:"type"` 
	Channel_id string `json:"channel_id"` 
	Url string `json:"url"` 
	Image_hash string `json:"image_hash"` 
	Position string `json:"position"` 
	Status string `json:"status"` 
	Created_at string `json:"created_at"` 
	Updated_at string `json:"updated_at"` 
	Package_id string `json:"package_id"` 
}

func get_promotion() {

    url := "http://www.miii.tv/channels/promotion/block"

    res, err := http.Get(url)

    if err != nil {
        panic(err.Error())
    }
    // fmt.Println(res.Body)
    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }
    // body := res.Body
    
    // body = []byte(`{"status": "ok","data": [{"id": "398","country_code": "TW","title": "服務終止","duration_from": "0000-00-00 00:00:00","duration_to": "0000-00-00 00:00:00","type": "1","channel_id": null,"url": "http://www.miii.tv/doc/eos","image_hash": "1438662919","position": "0","status": "1","created_at": "2015-08-04 12:34:22","updated_at": "2015-08-04 12:35:19","package_id": null},{"id": "397","country_code": "TW","title": "服務終止","duration_from": "0000-00-00 00:00:00","duration_to": "0000-00-00 00:00:00","type": "1","channel_id": null,"url": "http://www.miii.tv/doc/eos","image_hash": "1438662933","position": "0","status": "1","created_at": "2015-08-04 12:26:16","updated_at": "2015-08-04 12:35:33","package_id": null},{"id": "396","country_code": "TW","title": "服務終止","duration_from": "0000-00-00 00:00:00","duration_to": "0000-00-00 00:00:00","type": "1","channel_id": null,"url": "http://www.miii.tv/doc/eos","image_hash": "1438662943","position": "0","status": "1","created_at": "2015-08-04 12:19:28","updated_at": "2015-08-04 12:35:43","package_id": null}]}`)
    
    var data promotion_block
	json.Unmarshal(body, &data)
	// fmt.Println(data)
	fmt.Println("==============================")
	fmt.Println(data.Data[0])
	fmt.Println(data.Data[0].Id)
}



func get_list() {

    url := "http://www.miii.tv/channels/list/self"

    res, err := http.Get(url)

    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }
    
    // body = []byte(`{"status": "ok","data": [{"official": [{"ch_id": "57","ch_name": "UDN直播新聞台","ch_description": "","owner_id": "14","owner_name": "Ayukawa","owner_photo": "38/bc/14.jpg?1415009271","ch_number": "10","ch_type": "2","ch_rating_system": "0","ch_views": "10360","ch_likes": "0","ch_subscribes": "211","ch_has_cover": "1","ch_cover_hash": "1415107242"},{"ch_id": "1698","ch_name": "大愛一臺HD Live 直播","ch_description": "","owner_id": "618","owner_name": "Kelly","owner_photo": "c3/81/618.jpg?1423460588","ch_number": "11","ch_type": "2","ch_rating_system": "0","ch_views": "5639","ch_likes": "0","ch_subscribes": "72","ch_has_cover": "1","ch_cover_hash": "1431919773"}],"create": [ ],"subscribe": [ ]}]}`)
    
    var data list_self
	json.Unmarshal(body, &data)
	// fmt.Println(data)
	fmt.Println("==============================")
	fmt.Println(data.Data[0].Official[0])
	fmt.Println(data.Data[0].Official[0].Ch_id)
}

type list_self struct {
	Status string `json:"status"`
	Data []list_data `json:"data"`
}

type list_data struct {
	Official []official_array `json:"official"`
	Create []create_array `json:"create"`
	Subscribe []subscribe_array `json:"subscribe"`
}

type official_array struct {
	Ch_id string `json:"ch_id"`
	Ch_name string `json:"ch_name"`
	Ch_description string `json:"ch_description"`
	Owner_id string `json:"owner_id"`
	Owner_name string `json:"owner_name"`
	Owner_photo string `json:"owner_photo"`
	Ch_number string `json:"ch_number"`
	Ch_type string `json:"ch_type"`
	Ch_rating_system string `json:"ch_rating_system"`
	Ch_views string `json:"ch_views"`
	Ch_likes string `json:"ch_likes"`
	Ch_subscribes string `json:"ch_subscribes"`
	Ch_has_cover string `json:"ch_has_cover"`
	Ch_cover_hash string `json:"ch_cover_hash"`
}
type create_array struct {
}
type subscribe_array struct {
}

func main() {
	get_promotion()
	fmt.Println("------------------------------------------------------")
	get_list()
	fmt.Println("------------------------------------------------------")
	get_promotion_2()
	fmt.Println("------------------------------------------------------")
	get_promotion_3()
}

type promotion_block_2 struct {
	Status string `json:"status"`
	Data []struct {
		Id string `json:"id"` 
		Country_code string `json:"country_code"` 
		Title string `json:"title"` 
		Duration_from string `json:"duration_from"` 
		Duration_to string `json:"duration_to"` 
		Types string `json:"type"` 
		Channel_id string `json:"channel_id"` 
		Url string `json:"url"` 
		Image_hash string `json:"image_hash"` 
		Position string `json:"position"` 
		Status string `json:"status"` 
		Created_at string `json:"created_at"` 
		Updated_at string `json:"updated_at"` 
		Package_id string `json:"package_id"` 
	} `json:"data"`
}

func get_promotion_2() {

    url := "http://www.miii.tv/channels/promotion/block"

    res, err := http.Get(url)

    if err != nil {
        panic(err.Error())
    }
    fmt.Println(res.Body)
    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }
    // body := res.Body
    
    // body = []byte(`{"status": "ok","data": [{"id": "398","country_code": "TW","title": "服務終止","duration_from": "0000-00-00 00:00:00","duration_to": "0000-00-00 00:00:00","type": "1","channel_id": null,"url": "http://www.miii.tv/doc/eos","image_hash": "1438662919","position": "0","status": "1","created_at": "2015-08-04 12:34:22","updated_at": "2015-08-04 12:35:19","package_id": null},{"id": "397","country_code": "TW","title": "服務終止","duration_from": "0000-00-00 00:00:00","duration_to": "0000-00-00 00:00:00","type": "1","channel_id": null,"url": "http://www.miii.tv/doc/eos","image_hash": "1438662933","position": "0","status": "1","created_at": "2015-08-04 12:26:16","updated_at": "2015-08-04 12:35:33","package_id": null},{"id": "396","country_code": "TW","title": "服務終止","duration_from": "0000-00-00 00:00:00","duration_to": "0000-00-00 00:00:00","type": "1","channel_id": null,"url": "http://www.miii.tv/doc/eos","image_hash": "1438662943","position": "0","status": "1","created_at": "2015-08-04 12:19:28","updated_at": "2015-08-04 12:35:43","package_id": null}]}`)
    
    var data promotion_block_2
	json.Unmarshal(body, &data)
	// fmt.Println(data)
	fmt.Println("==============================")
	fmt.Println(data.Data[0])
	fmt.Println(data.Data[0].Id)
}


type promotion_block_3 struct {
	Status t3 `json:"status"`
}
type t3 map[string]string

func get_promotion_3() {

    url := "http://www.miii.tv/channels/promotion/block"

    res, err := http.Get(url)

    if err != nil {
        panic(err.Error())
    }
    // fmt.Println(res.Body)
    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }
    // body := res.Body
    
    body = []byte(`{"status": {"ok-k":"ok-v", "f-k": "f-v"}}`)
    
    var data promotion_block_3
	json.Unmarshal(body, &data)
	// for k, v := range data.Status {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }
	fmt.Println(data)
	fmt.Println("==============================")
	// fmt.Println(data.Status["ok-k"])
	// fmt.Println(data.Data[0].Id)
}