package main

import (
	"fmt"
)

func main() {

	// fmt.Println("Enter you number")
	// var num int
	// fmt.Scanf("%d", &num)

	// if num%2 == 0 {
	// 	fmt.Println("The number is even")
	// } else {
	// 	fmt.Println("The number is odd")
	// }

	var isAuthorized, isAuthenicated bool = false, false;

	if !isAuthenicated {
		fmt.Println("User is not Authenicated")
	} else if !isAuthorized {
		fmt.Println("User is not Authorized")
	} else{
		fmt.Println("User is good to go")
	}

}
