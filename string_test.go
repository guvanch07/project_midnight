package main

import "fmt"

func fStrings() {
	var a string
	fmt.Scan(&a)
	for i, v := range a {
		if i == len(a)-1 {
			fmt.Printf("%s", string(v))
			break

		}
		fmt.Printf("%s*", string(v))
	}

}
