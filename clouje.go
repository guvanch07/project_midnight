package main

import "fmt"

func clourje() {
	fn := func(x uint) (y uint) {
		for k := uint(1); x > 0; x /= 10 {
			if d := x % 10; d%2 == 0 && d != 0 {
				y += k * d
				k *= 10
			}
		}
		if y == 0 {
			y = 100
		}
		return
	}
	fmt.Println(fn(727178))
}
