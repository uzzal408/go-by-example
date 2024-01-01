package main

type person struct {
	name string
	age  int
}

func structFunc(name string, age int) person {

	p := person{name: name}
	p.age = age

	return p
}
