package main

import "fmt"

type rect struct {
	height, width int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) peri() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{height: 5, width: 10}

	fmt.Println("area: ", r.area())
	fmt.Println("peri: ", r.peri())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("peri: ", rp.peri())
}
