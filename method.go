package main

type rect struct {
	width, hieght float64
}

func (r *rect) area() float64 {
	return r.width * r.hieght
}

func (r *rect) perim() float64 {
	return 2*r.width + 2*r.hieght
}
