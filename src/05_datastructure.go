package main

import (
	"fmt"
)

func main() {
	//array
	var x [5]int
	x[3] = 43
	//x[5] = 50
	fmt.Println(x)

	//slices
	y := []int{3, 4, 5, 10}
	fmt.Println(y)

	fmt.Println(len(y))
	fmt.Println(cap(y))

	//loop through slice
	for i, v := range y {
		fmt.Println(i, v)
	}

	//Slicing the slices
	fmt.Println(y[:2])

	//Append to slices
	y = append(y, 12, 10, 14)
	//	fmt.Println(append(y, 11, 12, 14))
	fmt.Println(y)

	z := []int{20, 21}

	fmt.Println(append(y, z...))

	//Make
	w := make([]int, 10, 100)
	fmt.Println(w)
	fmt.Println(len(w))
	fmt.Println(cap(w))

	xp := [][]string{{"james", "bond"}, {"need", "of", "the", "hour"}}
	fmt.Println(xp)
}
