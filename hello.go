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
}
