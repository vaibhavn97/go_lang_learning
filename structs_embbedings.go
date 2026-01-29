package main

import "fmt"

type salary struct {
	base      int
	flexible  int
	allowance int
}

func (s salary) getGrossSalary() int {
	return s.allowance + s.base + s.flexible
}

type employee struct {
	id int
	salary
}

func main() {

	emp := employee{
		id: 1868,
		salary: salary{
			base:      100,
			flexible:  30,
			allowance: 20,
		},
	}

	fmt.Println(emp)
	fmt.Println(emp.getGrossSalary())

}
