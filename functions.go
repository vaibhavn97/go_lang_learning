package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func countNumber(nums []int, target int) ([]int, int) {
	cnt := 0
	for _, num := range nums {
		if num == target {
			cnt++
		}
	}
	return nums, cnt

}

func mulMany(nums ...int) int {
	res := 1
	for _, num := range nums {
		res *= num
	}
	return res
}



func main() {

	fmt.Println("Normal Function", add(1, 70))
	fmt.Println("Traling arguments", sub(1, 71))
	nums, count := countNumber([]int{4, 2, 3, 4}, 4)
	fmt.Println("Mutliple Return values", nums, count)
	fmt.Println("Variadic Functions", mulMany(1, 2, 3, 4))

	div := func (a, b int) int {
		return a/b
	}

	fmt.Println("Anonymous function", div(5, 3))
}
