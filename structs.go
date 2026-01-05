package main

import "fmt"

type person struct {
	name   string
	salary int
}

func (p *person) incrementSalary() {
	p.salary++;
}

func main() {

	per := person{"Arjun", 100}

	fmt.Println(per)

	cart := struct {
		id    int
		name  string
		price float32
	}{
		1,
		"book",
		12.21,
	}

	fmt.Println(cart.price)

	fmt.Println("Person salary", per.salary)
	per.incrementSalary()
	fmt.Println("Person incremented salary", per.salary)

}
