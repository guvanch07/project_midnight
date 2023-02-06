package main

import (
	"fmt"
	"math"
	"time"
)

var k, p, v float64 = 1296, 6, 6

func T() float64 {
	return 6 / W()
}
func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v

}

func findmap() {

	m := make(map[int]int)
	workArray := [10]int{}

	for i := range workArray {
		fmt.Scan(&workArray[i])
	}

	for i, u := range workArray {
		if v, ok := m[i]; ok {
			fmt.Print(v)
			fmt.Printf("%d ", v)

		} else {
			m[i] = work(u)
			fmt.Printf("%d ", m[i])
			return
		}
	}

}

func work(x int) int {
	time.Sleep(time.Second)
	if x >= 4 {
		return x + 1
	} else {
		return x - 1
	}
}

func findPopulation() {
	groupCity := map[int][]string{
		10:   {"Село", "Деревня", "ПГТ"},  // города с населением 10-99 тыс. человек
		100:  {"Город", "Станица"},        // города с населением 100-999 тыс. человек
		1000: {"Мегаполис", "Миллионник"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Город":     101,
		"Станица":   102,
		"Село":      103,
		"Мегаполис": 104,
	}

	for s := range cityPopulation {
		isOk := false
		for _, s2 := range groupCity[100] {
			if s2 == s {
				isOk = true
				break
			}
		}
		fmt.Println(isOk)
		if !isOk {
			delete(cityPopulation, s)
		}
	}
	fmt.Println(cityPopulation)
}

func findPopulation2() {
	groupCity := map[int][]string{
		10:   {"Село", "Деревня", "ПГТ"},  // города с населением 10-99 тыс. человек
		100:  {"Город", "Станица"},        // города с населением 100-999 тыс. человек
		1000: {"Мегаполис", "Миллионник"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Город":     101,
		"Станица":   102,
		"Село":      103,
		"Мегаполис": 104,
	}
	for _, city := range append(groupCity[10], groupCity[1000]...) {
		delete(cityPopulation, city)
	}

	for _, city := range groupCity[1000] {
		delete(cityPopulation, city)
	}

	for _, city := range groupCity[10] {
		delete(cityPopulation, city)
	}
}
