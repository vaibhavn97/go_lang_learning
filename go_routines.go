package main

import (
	"fmt"
	"time"
)

func getMemeberDetails(id int) {
	fmt.Println("Fetching for member", id)
	time.Sleep(2 * time.Second)
	fmt.Println("Fetched for member", id)
}

func main() {
	fmt.Println("Using Normal Function Call")
	getMemeberDetails(1)
	getMemeberDetails(2)
	getMemeberDetails(3)

	fmt.Println("Using Normal Function Call")
	go getMemeberDetails(10)
	go getMemeberDetails(20)
	go getMemeberDetails(30)

	time.Sleep(3 * time.Second)
}
