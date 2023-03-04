package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ResponseData []ResponseDatum

func UnmarshalResponseData(data []byte) (ResponseData, error) {
	var r ResponseData
	err := json.Unmarshal(data, &r)
	return r, err
}

type ResponseDatum struct {
	GlobalID       int     `json:"global_id"`
	SystemObjectID string  `json:"system_object_id"`
	SignatureDate  string  `json:"signature_date"`
	Razdel         Razdel  `json:"Razdel"`
	Kod            *string `json:"Kod,omitempty"`
	Name           string  `json:"Name"`
	Idx            string  `json:"Idx"`
	Nomdescr       *string `json:"Nomdescr,omitempty"`
}

type Razdel string

type (
	Student struct {
		Rating []int `json:"Rating"`
	}

	Group struct {
		Students []Student `json:"Students"`
	}

	Rating struct {
		Average float32
	}
)

func countAvarageRating() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	var group Group
	json.Unmarshal(data, &group)

	if err != nil {
		fmt.Println(err)
		return
	}
	var a, b float32
	for i := 0; i < len(group.Students); i++ {
		a++
		rating := group.Students[i].Rating
		for j := 0; j < len(rating); j++ {
			b++
		}
	}
	sum := b / a
	average := Rating{
		Average: sum,
	}
	indent, err := json.MarshalIndent(average, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", indent)
}

func fetchLocaly() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	res, err := UnmarshalResponseData(data)

	if err != nil {
		fmt.Println(err)
	}
	sum := 0
	for i := 0; i < len(res); i++ {
		sum += int(res[i].GlobalID)
	}
	fmt.Println(sum)
}

func fromNetwork() {
	var StructData ResponseData
	var urlData = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"
	sum := 0
	resp, err := http.Get(urlData)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	r := json.NewDecoder(resp.Body)
	r.Decode(&StructData)
	for _, val := range StructData {
		sum += val.GlobalID
	}
	fmt.Println("Result:", sum)

}
