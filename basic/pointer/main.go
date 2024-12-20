package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	c.Count++
}

func increment(x *int) {
	*x++
}

func rename(p *Person, name string) {
	p.Name = name
}

func main() {

	i := 42
	p := &i

	fmt.Println(*p)

	increment(p)

	fmt.Println(i)

	person := Person{Name: "John", Age: 30}
	rename(&person, "Jane")
	fmt.Println(person)

	// nil pointer => pointer that doesn't point to any memory address
	var j *int

	if j == nil {
		fmt.Println("j is nil")
	}

}
