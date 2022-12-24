package main

import (
	"fmt"
	"strconv"
)

func binaryFind() {
	var n int
	fmt.Scan(&n)
	i := int64(n)
	fmt.Println(strconv.FormatInt(i, 2))
}
