package main

import "fmt"

func mappers() {
	m := make(map[string]string)

	m["hey"] = "privet"
	m["how are you"] = "kak dela"
	fmt.Println(m)
}

func findhowmuchmatch() {
	// предложение
	sentence := "Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου"
	// инициализируем map
	// ключами будут знаки, а значениями — количество вхождений
	frequency := make(map[rune]int)
	// пройдёмся по знакам в предложении
	for _, v := range sentence {
		frequency[v]++ // встреченному знаку увеличиваем частоту
	}
	// печатаем
	for k, v := range frequency {
		fmt.Printf("Знак %c встречается %d раз \n", k, v)
	}
}

func addMoreMoneyToproduts() {
	pricelist := map[string]int{"хлеб": 50, "молоко": 100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}

	total := 0

	fmt.Println("Перечень деликатесов:")
	for k, v := range pricelist {
		if v > 500 {
			fmt.Println(k)
		}
	}
	for _, v := range order {
		total += pricelist[v]
	}

	fmt.Println("Стоимость заказа ", total)
}

func search(arr []int, k int) []int {
	// Создадим пустую map
	m := make(map[int]int)
	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {
		if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
			fmt.Println([]int{i, j})
			return []int{i, j}
		}
		// если искомого значения нет, то добавляем текущий индекс и значение в map
		m[a] = i
	}
	// не нашли пары подходящих чисел
	return nil
}

func removeduplicate(input []string) []string {

	output := make([]string, len(input))
	copy(output, input)

	inputSet := make(map[string]struct{}, len(input))
	outputIdx := 0
	for _, v := range input {
		if _, ok := inputSet[v]; !ok {
			output[outputIdx] = v
			outputIdx++
		}
		inputSet[v] = struct{}{}
	}
	fmt.Println(output[:outputIdx])
	return output[:outputIdx]

}
