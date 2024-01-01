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
}
