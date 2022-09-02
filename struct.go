package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type PersonModel struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func callStructToSerilize() {

	man := PersonModel{
		Name:        "Alex",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(),
	}
	jsMan, err := json.Marshal(man)

	if err != nil {
		log.Fatalln("unable marshal to json")
	}
	fmt.Printf("Man %v", string(jsMan))
}

// /task 2
const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}
`

type (
	Response struct {
		Header ResponseHeader `json:"header"`
		Data   ResponseData   `json:"data,omitempty"`
	}

	ResponseHeader struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
	}

	ResponseData []ResponseDataItem

	ResponseDataItem struct {
		Type       string                `json:"type"`
		Id         int                   `json:"id"`
		Attributes ResponseDataItemAttrs `json:"attributes"`
	}

	ResponseDataItemAttrs struct {
		Email      string `json:"email"`
		ArticleIds []int  `json:"article_ids"`
	}
)

func ReadResponse(rawResp string) (Response, error) {
	resp := Response{}
	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
	}

	return resp, nil
}

func callResponse() {
	resp, err := ReadResponse(rawResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp)
}

// /////////////////////////////////////////////
type PersonStruct struct {
	// Имя
	Name string
	// Количество детей
	Email string
	// Возраст
	dateOfBirth time.Time
}

// /создание конструктора
func newPersonStruct(name, email string, dobYear, dobMonth, dobDate int) PersonStruct {
	return PersonStruct{
		Name:        name,
		Email:       email,
		dateOfBirth: time.Date(dobYear, time.Month(dobMonth), dobDate, 0, 0, 0, 0, time.UTC),
	}
}

// func NewPersonStruct(name, email string, dobYear, dobMonth, dobDate int) (PersonStruct, error) {
// }

// ананимнав структура
func anonimousStruct() {
	req := struct {
		NameContains string `json:"name_contains"`
		Offset       int    `json:"offset"`
		Limit        int    `json:"limit"`
	}{
		NameContains: "Иван",
		Limit:        50,
	}
	reqRaw, _ := json.Marshal(req)
	fmt.Println(string(reqRaw))
}
