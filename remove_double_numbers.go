package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rmDoubledNum() {
	var n, d int
	fmt.Scan(&n, &d)
	x := strconv.Itoa(n)
	y := strconv.Itoa(d)

	fmt.Println(strings.ReplaceAll(x, y, ""))
}

func good() {
	var a, b int
	fmt.Scan(&a, &b)
	result := 0
	i := 1
	for a > 0 {
		if a%10 != b {
			result += a % 10 * i
			fmt.Println(result)
			i *= 10
		}
		a /= 10
	}
	fmt.Println(result)
}
