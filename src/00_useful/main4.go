package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"


type promotion_block struct {
	Status string `json:"status"`
	Data [] struct {
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
    fmt.Println(res.Body)
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
	// fmt.Println(data.Data[0])
	fmt.Println(data.Data[0].Id)
}

func main() {
	get_promotion()
}



