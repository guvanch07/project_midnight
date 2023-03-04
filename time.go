package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func timeFinde() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	t, err := time.Parse("2006-01-02 15:04:05", scanner.Text())

	if err != nil {
		panic(err)
	}
	if t.Hour() > 13 {
		added := t.AddDate(0, 0, 1)
		fmt.Println(added.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println(t.Format("2006-01-02 15:04:05"))
	}
}
