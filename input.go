package main

import "fmt"

func main() {

	fmt.Println("Give me a number")
	var num int
	fmt.Scanf("%d", &num)
	fmt.Print("You have typed ");
	fmt.Println(num);

}
