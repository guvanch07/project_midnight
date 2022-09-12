package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// we can dive more integers
func Sum(x ...int) (res int) {
	for _, v := range x {
		res += v
	}
	return
}

//bar(foo())

// func PrintAllFilesWithFilter(path string, filter string) {
// 	// получаем список всех элементов в папке (и файлов, и директорий)
// 	files, err := ioutil.ReadDir(path)
// 	if err != nil {
// 		fmt.Println("unable to get list of files", err)
// 		return
// 	}
// 	//  проходим по списку
// 	for _, f := range files {
// 		// получаем имя элемента
// 		// filepath.Join — функция, которая собирает путь к элементу с разделителями
// 		filename := filepath.Join(path, f.Name())
// 		// печатаем имя элемента, если путь к нему содержит filter
// 		if strings.Contains(filename, filter) {
// 			fmt.Println(filename)
// 		}
// 		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
// 		if f.IsDir() {
// 			PrintAllFilesWithFilter(filename, filter)
// 		}
// 	}
// }

func countCall(f func(string)) func(string) {
	cnt := 0
	// получаем имя функции. Подробнее об этом вызове будет рассказано в следующем курсе
	funcname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return func(s string) {
		cnt++
		fmt.Printf("Функция %s вызвана %d раз\n", funcname, cnt)
		f(s)
	}
}

// metricTimeCall — функция-обёртка для замера времени
func metricTimeCall(f func(string)) func(string) {
	return func(s string) {
		start := time.Now() // получаем текущее время
		f(s)
		fmt.Println("Execution time: ", time.Now().Sub(start)) // получаем интервал времени как разницу между двумя временными метками
	}
}

func myprint(s string) {
	fmt.Println(s)
}

func callAllClouse() {
	countedPrint := countCall(myprint)
	countedPrint("Hello world")
	countedPrint("Hi")

	// обратите внимание, что мы оборачиваем уже обёрнутую функцию, а значение счётчика вызовов при этом сохраняется
	countAndMetricPrint := metricTimeCall(countedPrint)
	countAndMetricPrint("Hello")
	countAndMetricPrint("World")
}
