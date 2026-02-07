package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	p := Person{
		Name:        "Alex",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC),
	}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("serialization error")
		return
	}

	str := string(data)

	fmt.Println(str)
}
