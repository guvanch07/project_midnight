package main

import (
	// пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func mjhjgj() {
	//value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
}

func readTask() (value1, value2, operation interface{}) {
	var val1, val2, oper interface{}

	if v, ok := value1.(float64); ok {
		val1 = v
	} else {
		fmt.Println("неизвестная операция")
		return
	}
	if v, ok := value2.(float64); ok {

		val2 = v
	} else {
		fmt.Println("неизвестная операция")
		return
	}

	return val1, val2, oper
}
