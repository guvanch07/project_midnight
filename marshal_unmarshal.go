package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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
