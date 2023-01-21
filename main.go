package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	exemaplerune()
}

func testStepik() {

	var a int
	var b int
	fmt.Scan(&a, &b)

	j := 0
	d := 0

	for i := 0; i <= a; i++ {

	}
	fmt.Print(j, d)

}

func m() {
	var x int
	var p int
	var y int

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	k := 0
	x = x * 100
	y = y * 100
	p = p + 100

	for x < y {
		x = (x * p) / 100
		k++
	}
	fmt.Println(k)
}

func gocal() {
	var a float64

	fmt.Scan(&a)

	if a <= 0 {
		fmt.Printf("число %2.2f не подходит\n", a)
	} else if a >= 10000 {
		fmt.Printf("%e\n", a)
	} else {
		doub := math.Pow(a, 2)
		fmt.Printf("%.4f\n", doub)
	}
}

func writeAllNum() {
	var n int
	fmt.Scan(&n)
	slice := make([]int, n)
	var sliceOdd []int

	if n >= 0 && n <= 100 {
		for i := range slice {
			_, _ = fmt.Scan(&slice[i])
			if i%2 == 0 {
				sliceOdd = append(sliceOdd, slice[i])
			}
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(sliceOdd), "[]"))
}

func newSl() {
	var n int
	fmt.Scan(&n)

	m := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m[i])
	}

	for i := 0; i < n; i += 2 {
		fmt.Print(m[i], " ")
	}
}

func findCount() {
	var n int
	count := 1
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a <= 0 {
			count++
		}
	}
	fmt.Println(count)
}
