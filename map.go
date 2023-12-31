package main

import (
	"fmt"
)

func mapFunc() {
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println("print map: ", m)

	fmt.Println("get map: ", m["c"])

	//delete form map
	delete(m, "a")
	fmt.Println("Deleted: ", m)

	_, prs := m["c"]
	fmt.Println("prs:", prs)

}
