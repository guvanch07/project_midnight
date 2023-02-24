package main

import (
	"fmt"
	"strconv"
)

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
func deleteOdd(n uint) uint {
	s := strconv.Itoa(int(n))
	var c string
	for _, v := range s {
		d, err := strconv.Atoi(string(v))
		if err != nil {
			panic(d)
		}
		if d%2 == 0 {
			if d > 0 {
				str := strconv.Itoa(d)
				c += str
			}
		}
	}
	result, _ := strconv.Atoi(c)
	fmt.Println(result)
	if result == 0 {
		return 100
	}
	return uint(result)
}
