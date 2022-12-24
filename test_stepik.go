package main

import (
	"fmt"
	"strconv"
)

func stepik() {
	var a int
	fmt.Scan(&a)

	h := a / 30
	m := 2 * (a % 30)

	result := "It is " + strconv.Itoa(int(h)) + "hours " + strconv.Itoa(int(m)) + "minutes"

	fmt.Println(result)
}

func timeFinder() {
	var input int
	fmt.Scan(&input)

	h := input / 3600
	m := (input % 3600) / 60
	fmt.Printf("It is %o hours %d minutes.\n", h, m)
}
