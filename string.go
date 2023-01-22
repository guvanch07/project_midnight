package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func exemaplerune() {
	var x string
	fmt.Scan(&x)
	for i, char := range x {
		if i%2 != 0 {
			fmt.Print(string(char))
		}
	}
}

func findCountofString() {
	var s string
	fmt.Scan(&s)
	for _, v := range s {
		c := strings.Count(s, string(v))
		if c == 1 {
			fmt.Print(string(v))
		}
	}
}

func checkPassword() bool {
	var s string
	fmt.Scan(&s)
	reg := regexp.MustCompile(`^[a-zA-Z0-9]{5,}$`).MatchString
	return reg(s)
}

func countAllString() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rn := []rune(text)
	b := unicode.IsUpper(rn[0])
	l := rn[len(rn)-2] == '.'
	fmt.Println(b)
	fmt.Println(l)
	fmt.Println(len(rn))
	if b && l {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

//быть или не быть.
