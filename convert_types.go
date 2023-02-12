package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertor() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(s)
	s1 := strings.Split(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), ",", "."), ";")

	res, e := strconv.ParseFloat(s1[0], 64)
	if e != nil {
		panic(e)
	}
	res1, e1 := strconv.ParseFloat(s1[1], 64)
	if e != nil {
		panic(e1)
	}

	fmt.Printf("%.4f", res/res1)

}

func anon() {
	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Printf("%T", x)
}
