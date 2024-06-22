package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 24})
	fmt.Println(&person{name: "Ann", age: 56})
	fmt.Println(newPerson("John")) // just name, no age as input

	s := person{name: "Jim", age: 10}
	sp := &s
	fmt.Println(sp.age)

	sp.age = 50
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
}
