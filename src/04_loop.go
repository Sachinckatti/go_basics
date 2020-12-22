package main

import (
	"fmt"
)

func main() {
	//Control flow
	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}

	//Conditional statement
	x := 20
	if x < 10 {
		fmt.Println(x, "is less than", 10)
	} else {
		fmt.Println(x, " is greater than ", 10)
	}

	if x == 5 {
		fmt.Println(x, " is equal to ", 5)
	}
}
