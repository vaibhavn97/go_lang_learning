package main

import "fmt"

func main() {
	fmt.Println("Array")
	nums := [3]int{1, 2, 3}
	fmt.Println(nums)
	nums[1] = 80
	// range - index, element
	for i := range nums {
		fmt.Print(nums[i])
	}
	fmt.Println();
	for _, num := range nums {
		fmt.Print(num)
	}

	// Cannot add elements

}
