package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func file() {
	dataFor := []byte("123")
	ioutil.WriteFile("test.txt", dataFor, 0600)
	ioutil.WriteFile("test.txt", dataFor, 0600)
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}

func fileOS() {
	for i := 0; i <= 3; i++ {
		file, err := os.Create(strconv.Itoa(i) + "text.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
	os.Rename("2text.txt", "4text.txt")
	for i := 4; i >= 2; i-- {
		os.Remove(strconv.Itoa(i) + "text.txt")
	}
}

func fileOS1() {
	var f *os.File
	file1, _ := os.Create("text.txt")
	file1.WriteString("1")
	file1.WriteString("3")
	switch file1.Name() {
	case "text.txt":
		os.Rename("text.txt", "textxt.txt")
		fallthrough
	case "texttt.text":
		os.Remove("text.txt")
		fallthrough
	default:
		var err error
		f, err = os.Open("textxt.text")
		file1.WriteString("2")
		if err != nil {
			panic(err)
		}

	}
	res, _ := ioutil.ReadAll(f)
	resInt, err := strconv.Atoi(string(res[2]))
	if err != nil {
		panic(err)
	}
	fmt.Println(resInt + 1)
	file1.Close()
}
