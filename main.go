package main

import "fmt"

func main() {
	coralSlice := []string{"blue coral", "foliose coral", "pillar coral", "elkhorn coral", "black coral", "antipathes", "leptopsammia", "massive coral", "soft coral"}

	coralSlice = append(coralSlice[:1], coralSlice[8:]...)

	fmt.Printf("%q\n", coralSlice)
}
