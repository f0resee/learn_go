package main

import "fmt"

type Color interface {
	Clone() Color
	Print() string
}

type Red struct {
	color string
}

func (r *Red) Clone() Color {
	return &Red{
		color: r.color,
	}
}

func (r *Red) Print() string {
	return r.color
}

func main() {
	r0 := Red{
		color: "red",
	}
	fmt.Println(r0.Print())

	r1 := r0.Clone()
	fmt.Println(r1.Print())
}
