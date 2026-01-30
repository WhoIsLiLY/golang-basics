package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.Name, p.Age)
}

func (p *Person) Birthday() {
	p.Age++
	fmt.Printf("%s is now %d years old\n", p.Name, p.Age)
}

func main() {
	person := Person{
		Name: "David",
		Age:  25,
	}

	person.Introduce()
	person.Birthday()
	person.Introduce()
}