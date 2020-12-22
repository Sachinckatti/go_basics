package main

import (
	"fmt"
)

func main() {
	mp := map[string]int{}
	mp["james"] = 4
	mp["bond"] = 20

	fmt.Println(mp)

	if v, ok := mp["banana"]; ok {
		fmt.Println("Banana is present", v)
	} else {
		fmt.Println("Banana is not present")
	}
}
