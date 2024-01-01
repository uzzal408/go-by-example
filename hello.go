package main

import "fmt"

func main() {
	fmt.Println("Hello! World")
	value()
	variable()
	array()
	slice()
	mapFunc()

	rangeFunc()
	value1, value2 := multipleReturn()
	fmt.Println("return multiple value: ", value1, value2)
	sum(10, 5, 4, 7)

	nums := []int{8, 9, 4}
	sum(nums...)

	nexItem := intSeq()
	fmt.Println(nexItem())
	fmt.Println(nexItem())
	fmt.Println(nexItem())

	newItem := intSeq()
	fmt.Println(newItem())

	person := structFunc("Ismail", 29)

	fmt.Println(person.name)

	r := rect{width: 20, hieght: 30}

	fmt.Println("area: ", r.area())
	fmt.Println("Perim: ", r.perim())

	rp := &r

	fmt.Println("RP Area: ", rp.area())
	fmt.Println("RP Perim: ", rp.perim())

}
