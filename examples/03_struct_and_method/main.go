package main

import "fmt"

func main() {
	fmt.Println("Person Example")
	p := Person{
		Name: "John",
		Age:  30,
	}

	p.Greet()
	p.SetName("Jane")
	p.Greet()

	fmt.Println("Point Example")
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 4, Y: 6}
	fmt.Println(p1.DistanceTo(p2))
	p1.Move(1, 1)
	fmt.Println(p1)

	fmt.Println("Interface Example")
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}
	shapes := []Shape{c, r}
	printArea(shapes)
}
