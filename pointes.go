package main

import "fmt"

type A struct {
	IntField int
}

// Литерал А{} создаёт в памяти переменную типа А. Затем от неё берётся указатель

func pointers() {
	p := &A{
		IntField: 10,
	}

	po := &A{}
	po.IntField = 42

	fmt.Println(p)
}

func anotherWay() {
	//or
	p := new(A)

	fmt.Println(p)
}

func readPonters() {
	i := 32
	p := &i //  save adress

	fmt.Println(*p) //read the adress
	*p = 21         // assign new value
}

func inrementTest() {
	incrementCopy := func(i int) {
		i++
	}

	increment := func(i *int) {
		(*i)++
	}

	i := 42

	incrementCopy(i)
	fmt.Println(i) // 42

	increment(&i)
	fmt.Println(i) // 43

}

func tackPointes() {
	a := 1
	p := &a
	b := &p
	*p = 3
	**b = 5
	a += 4 + *p + **b
	fmt.Printf("%d", *p)

}
