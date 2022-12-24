package main

import (
	. "fmt"
)

func theBigSeven() {
	var a, b int
	Scan(&a, &b)
	b = b / 7 * 7
	if b >= a {
		Println(b)
	} else {
		Println("NO")
	}
}

func korovaa() {
	var n int
	Scan(&n)

	if n > 0 && n < 100 {
		if n == 1 || n%10 == 1 {
			Println(n, "korova")
		} else if 2 <= n && n <= 4 {
			Println(n, "korovy")
		} else if 5 <= n && n <= 20 {
			Println(n, "korov")
		} else if n > 20 && (n%10 == 2 || n%10 == 3 || n%10 == 4) {
			Println(n, "korovy")
		} else {
			Println(n, "korov")
		}
	}

}

func findDoble() {
	var n int
	Scan(&n)
	for i := 1; i <= n; i += i { // i *= 2
		Print(i, " ")
	}
}
