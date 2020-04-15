package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Work() {
	fmt.Printf("%s is working\n", p.Name)
}
