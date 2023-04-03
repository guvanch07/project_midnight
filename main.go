package main

import (
	"fmt"
)

func greet(roc <-chan string) {
	fmt.Println("Hello " + <-roc + "!")
}

func main() {
	done := make(chan struct{})
	go func(c chan struct{}) {
		go work1(c)
		close(c)
	}(done)
}

func work1(c chan struct{}) {
	fmt.Println("hello")

}

// ///////////////////////////////////
func workIt(done chan struct{}) {
	fmt.Println("run work")
}

func callInputStream() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	go func() {
		defer close(inputStream)
		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()
	for x := range outputStream {
		fmt.Print(x)
	}
}

func removeDuplicates(inputStream, outputStream chan string) {
	m := make(map[string]string)
	for v := range inputStream {
		if m[v] != v {
			m[v] = v
			outputStream <- m[v]
		}
	}
	defer close(outputStream)
}

func task2(c chan string, n string) {
	for i := 0; i < 5; i++ {
		c <- n + " "
		fmt.Print(<-c)
	}
}
