package main

import "fmt"

func fibo() {
	var n int
	fmt.Scan(&n)
	var f []int

	is := checkfibonacci(n)

	if is {
		for i := 0; i <= n; i++ {
			f = append(f, fibonacciRecursion(i))
		}

		for i, v := range f {
			if v == n {
				fmt.Println(i)
			}
		}
	} else {
		fmt.Println(-1)
		return
	}

}

func fibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)

}

func checkfibonacci(n int) bool {
	a := 0
	b := 1
	if n == a || n == b {
		fmt.Println(true)
		return true
	}
	c := a + b
	for c <= n {
		if c == n {
			fmt.Println(true)
			return true
		}
		a = b
		b = c
		c = a + b
	}
	fmt.Println(false)
	return false
}
