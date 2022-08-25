package main

import "fmt"

func mappers() {
	m := make(map[string]string)

	m["hey"] = "privet"
	m["how are you"] = "kak dela"
	fmt.Println(m)
}
