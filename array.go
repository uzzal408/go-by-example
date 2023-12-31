package main

import "fmt"

func array() {
	var a [5]int
	fmt.Println("Array:", a)
	a[4] = 10
	fmt.Println("Array SET: ", a)
	fmt.Println("Array len: ", len(a))

	b := [4]int{1, 2, 3, 4}
	fmt.Println("Intial Set Array: ", b)

	//2 Dimention

	var twoD [3][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			twoD[i][j] = i + j
		}

	}

	fmt.Println("Print 2d array: ", twoD)

}
