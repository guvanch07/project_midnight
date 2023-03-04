package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func timeFormat() {
	var rf string
	fmt.Scan(&rf)
	formattedDate, err := time.Parse(time.RFC3339, rf)
	if err != nil {
		fmt.Println("Error while parsing the date time :", err)
	}
	fmt.Println(formattedDate.Format(time.UnixDate))
}

func taskTime4() {
	const now = 1589570165
	var min, sec int64
	fmt.Scanf("%d мин. %d", &min, &sec)
	fmt.Println(min, sec)

	minToSec := 60 * min
	sum := minToSec + sec + now
	t := time.Unix(sum, 0)
	fmt.Println(t.Format(time.UnixDate))
}

func taskTime3() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	sep := strings.Split(input.Text(), ",")
	fmt.Println(sep[1])

	t1, er1 := time.Parse("02.01.2006 15:04:05", sep[0])
	t2, er2 := time.Parse("02.01.2006 15:04:05", sep[1])
	fmt.Println(t1)
	fmt.Println(t2)
	if er1 != nil {
		panic(er1)
	}
	if er2 != nil {
		panic(er2)
	}
	t := time.Since(t2) - time.Since(t1)
	fmt.Println(t.Round(time.Second))
}
