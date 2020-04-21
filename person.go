package main

import "fmt"

func (p *Person) Work() {
	fmt.Printf("%s is working\n", p.Name)
}
