package main

import "fmt"

var V int

func forLoop() {
	v := 13
	for i := 0; i < 10; i++ {
		v++
		V = i
	}
	fmt.Println(v)
}

func rangeLoop() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	for arrayIndex, arrayValue := range array {
		fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
	}
}
