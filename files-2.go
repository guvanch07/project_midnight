package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func bufioReader() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		i, er := strconv.Atoi(scanner.Text())
		if er != nil {
			panic(er)
		}
		sum += i

	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
}

func cvsReader() {
	buf := bytes.NewBuffer(nil)

	w := csv.NewWriter(buf)

	for i := 1; i <= 3; i++ {
		// Запись данных может производится поэтапно, например в цикле
		val1 := fmt.Sprintf("row %d col 1", i)
		val2 := fmt.Sprintf("row %d col 2", i)
		val3 := fmt.Sprintf("row %d col 3", i)
		if err := w.Write([]string{val1, val2, val3}); err != nil { // Аргументом Write является срез строк
			// ...
		}
	}
	w.Flush() // Этот метод приведет к фактической записи данных из буфера csv.Writer в buf

	// Либо данные можно записать за один раз
	w.WriteAll([][]string{ // Аргументом WriteAll является срез срезов строк
		{"row 4 col 1", "row 4 col 2", "row 4 col 3"},
		{"row 5 col 1", "row 5 col 2", "row 5 col 3"},
	})

	r := csv.NewReader(buf)

	for i := 1; i <= 2; i++ {
		// Читать данные мы тоже можем построчно, получая срез строк за каждую итерацию
		row, err := r.Read()
		if err != nil && err != io.EOF { // Здесь тоже нужно учитывать конец файла
			// ...
		}
		fmt.Println(row)
	}

	// Либо прочитать данные за один раз
	data, err := r.ReadAll()
	if err != nil {
		// Когда мы читаем данные до конца файла io.EOF не возвращается, а служит сигналом к завершению чтения
		// ...
	}

	for _, row := range data {
		fmt.Println(row)
	}
}

func mywalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if strings.Contains(info.Name(), "co") {
		fmt.Printf("%d ", len(info.Name())+1)
		return nil
	} else {
		return nil
	}
}

func findInfoFile() {
	const root = "."
	if err := filepath.Walk(root, mywalkFunc); err != nil {
		fmt.Printf("ошибка: %v ", err)
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	buf, _ := os.Open(path)
	defer buf.Close()
	r := csv.NewReader(buf)
	row, err := r.ReadAll()
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}
	if len(row) > 1 {
		fmt.Println(row[4][2])
	}

	return nil
}

func findCvs42() {
	const root = "./task"
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}

func findZero() {
	file, err := os.Open("./task/nums.txt")
	if err != nil {
		return
	}
	defer file.Close()
	r := bufio.NewReader(file)
	var count int
	for {
		str, err := r.ReadString(';')
		if err != nil {
			break
		}
		count++

		if str == "0;" {
			fmt.Print(count)
			break
		}
	}
}
