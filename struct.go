package main

type person struct {
	name string
	age  int
}

func structFunc(name string, age int) *person {

	p := person{name: name, age: age}
	return &p
}
