package main

import "fmt"

func slice() {
	var a []string
	fmt.Println("empty Slice: ", a, len(a))
	s := make([]string, 4)
	s[0] = "1"
	s[1] = "2"
	s[2] = "3"
	s[3] = "4"
	fmt.Println("Assign Value in Slince: ", s)

	s = append(s, "c")
	s = append(s, "d")
	fmt.Println("append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	l := c[2:4]
	fmt.Println("part: ", l)

	twoD := make([][]int, 3)

	for k := 0; k < 3; k++ {
		innerLn := k + 1
		twoD[k] = make([]int, innerLn)
		for i := 0; i < innerLn; i++ {
			twoD[k][i] = k + i
		}
	}

	fmt.Println("2d Slice: ", twoD)
}
