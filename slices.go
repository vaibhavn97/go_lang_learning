package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums []int
	fmt.Println(nums)
	nums = append(nums, 1)
	fmt.Println(nums)

	stringArray := make([]string, 3)
	fmt.Println(stringArray)
	stringArray[0] = "123"
	stringArray[1] = "ABC"
	stringArray[2] = "9090"
	fmt.Println(stringArray)
	fmt.Println("Len is ", len(stringArray))

	stringArray = append(stringArray, "appended_element")
	fmt.Println(stringArray, len(stringArray))

	fmt.Println(stringArray[:2])

	if slices.Equal(nums, []int{1}) {
		fmt.Println("They are equal")
	} else {
		fmt.Println("Slices are not equal")
	}

}
