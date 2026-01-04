package main

import "fmt"

func nextNum() func() int {
	count := 0
	return func() int {
		count++
		return count;
	}
}

func main() {

	nextSeq := nextNum();

	fmt.Println(nextSeq());
	fmt.Println(nextSeq());
	fmt.Println(nextSeq());
	fmt.Println(nextSeq());
}
