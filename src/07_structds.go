package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type storeManager struct {
	person
	order bool
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(p2.first, p2.last)

	sm1 := storeManager{
		person: person{
			first: "Scare",
			last:  "Crow",
		},
		order: true,
	}

	fmt.Println(sm1)
	fmt.Println(sm1.first, sm1.last, sm1.order)
}
