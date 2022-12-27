package main

import (
	"math"
)

func fibonacciRec(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci(n int) int {
	first, second := 1, 1
	for i := 1; i < n; i++ {
		first, second = second, first+second
	}
	return first
}

func fibonacciMath(n int) int {
	if n < 1 {
		return 0
	}
	fib := math.Round(math.Pow((1+math.Sqrt(5))/2, float64(n)) / math.Sqrt(5))
	return int(fib)

}

func sumInt(a ...int) (int, int) {
	sum := 0
	for _, x := range a {
		sum += x
	}
	return len(a), sum
}
