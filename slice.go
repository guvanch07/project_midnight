package main

import (
	"fmt"
	"reflect"
)

// mySlice := make([]TypeOfelement, LenOfslice, CapOfSlice)
// mySlice := make([]int) // слайс [], базовый массив []
// mySlice := make([]int,5) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0]
// mySlice := make([]int,5,10) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0 0 0 0 0 0]

// s := make([]int, 4, 7) // [0 0 0 0], len = 4 cap = 7
// // 1. Создаём слайс s с базовым массивом на 7 элементов.
// // Четыре первых элемента будут доступны в слайсе.

// slice1 := append(s[:2], 2, 3, 4)
// fmt.Println(s, slice1) // [0 0 2 3] [0 0 2 3 4]
// // 2. Берём слайс из первых двух элементов s и добавляем к ним три элемента.
// // Так как суммарная длина полученного слайса (len == 5) меньше ёмкости s[:2] (cap == 7),
// // то базовый массив остаётся прежним.
// // Слайс s тоже изменился, но его длина осталась прежней

// slice2 := append(s[1:2], 7)
// fmt.Println(s, slice1, slice2) // [0 0 7 3] [0 0 7 3 4] [0 7]
// // 3. Здесь также базовый массив остаётся прежним, изменились все три слайса

// slice3 := append(s, slice1[1:]...)
// fmt.Println(len(slice3), cap(slice3))  // 8 14
// 4. Длина s и slice1[1:] равна 4, длина нового слайса будет равна 8,
// что больше ёмкости базового массива.
// Будет создан новый базовый массив ёмкостью 14,
// ёмкость нового базового массива подбирается автоматически
// и зависит от текущего размера и количества добавленных элементов

// // 5. Легко проверить, что slice3 ссылается на новый базовый массив
// s[1] = 99
// fmt.Println(s, slice1, slice2, slice3)
// [0 99 7 3] [0 99 7 3 4] [99 7] [0 0 7 3 0 7 3 4]

func copyMethod() {
	var dest []int
	dest2, dest3 := make([]int, 3), make([]int, 5)
	src := []int{1, 2, 3, 4}
	copy(dest, src)
	copy(dest2, src)
	copy(dest3, src)
	fmt.Println(dest, dest2, dest3, src) // [] [1 2 3] [1 2 3 4 0] [1 2 3 4]
}

func defendPanic() {
	s := []int{1, 2, 3}
	if len(s) != 0 { // защищаемся от паники
		s = s[:len(s)-1]
	}
	fmt.Println(s) // [1 2]
}

func removeFirst() {
	s := []int{1, 2, 3}
	if len(s) != 0 { // защищаемся от паники
		s = s[1:]
	}
	fmt.Println(s) // [2 3]
}
func removewithIndex() {
	s := []int{1, 2, 3, 4, 5}
	i := 2

	if len(s) != 0 && i < len(s)-1 { // защищаемся от паники
		s = append(s[:i], s[i+1:]...)
	}
	fmt.Println(s) // [1 2 4 5]
}

func compareSlices() {

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s3 := []string{"1", "2", "3"}
	s4 := []int{1, 2, 3}

	fmt.Println(reflect.DeepEqual(s1, s2)) // false
	fmt.Println(reflect.DeepEqual(s1, s3)) // false
	fmt.Println(reflect.DeepEqual(s1, s4)) // true
}

func tackSlice() {
	dim := 100
	s := make([]int, 0, dim)
	// заполняем слайс числами
	for i := 0; i < dim; i++ {
		s = append(s, i+1)
	}
	// оставляем первые и последние 10 элементов
	s = append(s[:10], s[dim-10:]...)
	//s = s[10 : dim-10]
	dim = len(s)
	//разворачиваем слайс
	for i := range s[:dim/2] {
		s[i], s[dim-i-1] = s[dim-i-1], s[i]
	}
	fmt.Println(s)
}

// Разернуть строку таким образом не получится, так как строка — неизменяемый тип данных.
// Строку можно преобразовать к слайсу байт ([]byte), но это опасно, если строка содержит Unicode-символы
// Лучше создать слайс рун из строки и развернуть его
// Например, так:

func revrseStringOfSlice() {
	input := "The quick brown 狐 jumped over the lazy 犬"
	// Get Unicode code points.
	n := 0
	// создаём слайс рун
	runes := make([]rune, len(input))
	// добавляем руны в слайс
	for _, r := range input {
		runes[n] = r
		n++
	}
	// убираем лишние нулевые руны
	runes = runes[0:n]
	// разворачиваем
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// преобразуем в строку UTF-8.
	output := string(runes)
	fmt.Println(output)
}

func badForPerformance() {
	numbers := []int{}
	for i := 0; i < 4; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)
}

// Хотя с функциональной точки зрения
// использование append() и cap() эквивалентно, в примере
// с cap() не выделяется лишняя память, которая потребовалась бы
// при использовании функции append().

func goodForPerformance() {
	numbers := make([]int, 4)
	for i := 0; i < cap(numbers); i++ {
		numbers[i] = i
	}

	fmt.Println(numbers)
}
