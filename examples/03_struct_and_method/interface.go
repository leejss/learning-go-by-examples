package main

import (
	"fmt"
	"math"
)

// Interface

// 인터페이스 정의
type Shape interface {
	Area() float64
}

func printArea(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
	}
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
