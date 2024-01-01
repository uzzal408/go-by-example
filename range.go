package main

import "fmt"

func rangeFunc() {
	m := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range m {
		sum += num
	}
	fmt.Println("Summation with range:", sum)

	for i, v := range m {
		if v == 3 {
			fmt.Println("Index of value: ", i)
		}
	}

	kvs := map[string]string{"a": "Apple", "b": "Ball", "c": "Cat"}

	for k, v := range kvs {
		fmt.Printf("%s-%s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
