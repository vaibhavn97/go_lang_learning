package main

import "fmt"

func main() {
	number := 10

	ptr := &number

	fmt.Println("Address of number", &number)
	fmt.Println("Content of pointer", ptr)
	fmt.Println("Address of pointer", &ptr)
	fmt.Println("Content of address hold by pointer", *ptr)

}
