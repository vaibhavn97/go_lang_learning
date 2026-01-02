package main

import "fmt"

func main() {
	fmt.Println("For Loops")

	fmt.Println("For loop iterating over numbers")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("For loop as while")
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

}
