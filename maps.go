package main

import "fmt"

func main() {

	nums := [4]int{1, 3, 4, 3}

	mp := make(map[int]int)

	for _, num := range nums {
		mp[num]++
	}

	fmt.Println(mp)
	fmt.Println("key value")
	for key, value := range mp {
		fmt.Println(key, " ", value)
	}

	delete(mp, 1)

	fmt.Println("After deleting 1", mp)


	clear(mp)

	fmt.Println("Clearing All keys", mp)

}
