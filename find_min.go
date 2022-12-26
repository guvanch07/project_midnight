package main

import "fmt"

func minimumFromFour() int {
	//print your code here
	var a, b, c, d int
	var slice []int
	fmt.Scan(&a, &b, &c, &d)
	slice = append(slice, a, b, c, d)
	min := slice[0]
	for _, v := range slice {
		if min > v {
			min = v
		}
	}

	fmt.Println(min)
	return min
}

func minimumFromFour2() int {
	//print your code here
	var res int
	fmt.Scan(&res)
	for i := 1; i < 4; i++ {
		var a int
		fmt.Scan(&a)
		if res > a {
			res = a
		}
	}
	return res
}

func minimumFromFour3() int {
	var min, num int
	for i := 0; i < 4; i++ {
		fmt.Scan(&num)
		if i == 0 || num < min {
			min = num
		}
	}
	return min
}
