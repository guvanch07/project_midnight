package main

import "fmt"

func whileLoop() {
	// создаём переменную
	i := 0
	// описываем предусловие
	for i < 5 {
		// наращиваем переменную
		i++
	}
	// выводим результат на экран
	fmt.Println(i)
}
