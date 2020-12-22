package main

import (
	"fmt"
)

func main() {

	foo()
	bar("phrase")
	s1 := woo()
	fmt.Println(s1)
}

func foo() {
	fmt.Println("Foo is called")
}

func bar(s string) {
	fmt.Println("Passed an argument ", s)
}

func woo() string {

	return "Pennyworth"
}
