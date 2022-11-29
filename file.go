package main

import "fmt"

func main() {
	fmt.Println("Hello World!!!")
	foo()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Index: ", i)
		}

	}
	boo()
}
func boo() {
	fmt.Println("Boo FUnction!!!!")
}
func foo() {
	fmt.Println("Welcome to Go Lang programming!!!!")
}
