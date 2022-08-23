package main

import "time"

type Person struct {
	Name        string
	Age         int
	lastVisited time.Time
}

func GetPersonWithLastVisisted(p Person) Person { //bad
	return Person{
		Name:        p.Name,
		Age:         p.Age,
		lastVisited: time.Now(), // time.Now() возвращает текущее время
	}
}

func callBad() {
	p := Person{
		Name:        "Alex",
		Age:         25,
		lastVisited: time.Time{}, // пустое значение времени — пользователь ещё не посещал наш сервис
	}

	p = GetPersonWithLastVisisted(p)
}

func UpdatePersonWithLastVisisted(p *Person) { //good
	p.lastVisited = time.Now()
}

func callGood() {
	p := Person{
		Name:        "Alex",
		Age:         25,
		lastVisited: time.Time{},
	}

	UpdatePersonWithLastVisisted(&p)
}
