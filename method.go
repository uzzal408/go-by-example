package main

type rect struct {
	width, hieght int
}

func (r *rect) area() int {
	return r.width * r.hieght
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.hieght
}
