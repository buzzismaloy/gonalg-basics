package main

import "fmt"

type Item struct {
	NoOption   string
	Parameter1 string
	Parameter2 int
}

// Here, options are functions applied to an object. For this, the approach is called funcopts.
func NewItem(opts ...option) *Item {
	i := &Item{
		NoOption:   "usual",
		Parameter1: "default",
		Parameter2: 42,
	}

	for _, opt := range opts {
		opt(i)
	}

	return i
}

type option func(*Item)

func Option1(option1 string) option {
	return func(i *Item) {
		i.Parameter1 = option1
	}
}

func Option2(option2 int) option {
	return func(i *Item) {
		i.Parameter2 = option2
	}
}

func main() {
	item1 := NewItem()
	item2 := NewItem(Option2(70))
	item3 := NewItem(Option1("rare"), Option2(99))
	item4 := NewItem(Option2(44), Option1("common"))

	fmt.Println(item1, item2, item3, item4)
}
