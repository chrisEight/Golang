package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()
	fmt.Println("foo printing is done")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			n, _ := fmt.Println(i)
			fmt.Println(n, "\n")
		}
	}
	bar()
}
func foo() {
	fmt.Println("I'm in foo")
}
func bar() {
	fmt.Println("BYEEEEEE BYEEEEE")
}
