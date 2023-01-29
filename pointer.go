package main

import "fmt"

func pointer() {
	a := 200
	b := &a
	*b++
	// b - 0xc00001a0b0
	// a - 201
	fmt.Println(a)
}

func one(xPt *int) {
	*xPt = 100 + 50
}
func pointOne() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(xPtr) // x is 1
}

func changer(i *int) {
	*i = 2 * *i
}
func changerMain() {
	i := 11
	changer(&i)
	fmt.Println(i)
}

func testPointerStar(x1 *int, x2 *int) {
	// a := *x2
	// *x2 = *x1
	// *x1 = a
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}

func PointerInPointer() {
	a := 200
	b := &a
	*b++
	c := &b
	**c++ // указатель на указатель
	fmt.Println(a)
}
