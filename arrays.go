package main

import "fmt"

var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}

func iteralarray() {
	sumTemp := 0
	for i := 0; i < len(weekTemp); i++ {
		sumTemp += weekTemp[i]
	}
	average := sumTemp / len(weekTemp)
	fmt.Println(average)
}

func gooditeralArray() {
	sumTemp := 0
	for _, temp := range weekTemp {
		sumTemp += temp
	}
	average := sumTemp / len(weekTemp)
	fmt.Println(average)
	// weekTemp [5 4 6 8 11 9 5] ! — значения не изменились

	for i := range weekTemp {
		weekTemp[i] = 0
	}
	// weekTemp [0 0 0 0 0 0 0] ! — значения изменились
}
