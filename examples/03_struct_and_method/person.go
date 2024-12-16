package main

import "fmt"

// collection of fields
type Person struct {
	Name string
	Age  int
}

// Value receiver
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// Pointer receiver
func (p *Person) SetName(name string) {
	p.Name = name
}
