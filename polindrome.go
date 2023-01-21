package main

import (
	"fmt"

	"regexp"
	"strings"
)

func isPolindrome1(str string) bool {
	bt := []rune(str)
	fmt.Println(bt)
	for i := len(bt) - 1; i >= 0; i-- {
		bt = append(bt, bt[i])
	}
	fmt.Println(string(bt))
	return str == string(bt)
}

func isPalindrome2(str string) bool {
	bt := []rune(str)
	for i := 0; i < len(bt); i++ {
		j := len(bt) - 1 - i
		if bt[i] != bt[j] {
			return false
		}
	}
	return true
}

func isPalindrome3(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	for i := range str {
		if str[i] != reversedStr[i] {
			return false
		}
	}
	return true
}

func palindrome(str string) bool {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")

	cleanstr := strings.ToLower(reg.ReplaceAllString(str, ""))
	arr := []rune(cleanstr)
	bnoden := []rune{}
	isPalindrome := false
	count := 0

	for i := len(arr) - 1; i >= 0; i-- {
		bnoden = append(bnoden, arr[i])
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == bnoden[i] {
			count++
		}
	}
	if count == len(arr) {
		isPalindrome = true
	}
	return isPalindrome
}
