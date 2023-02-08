package main

import "fmt"

func updateValue(x *string) {
	*x = "thakur"
}

func main() {
	name := "shivansh"

	m := &name
	fmt.Println("Address of name is ", m)
	fmt.Println("Value in name is", *m)

	updateValue(m)
	fmt.Println("Updated value using pointer is", name)
}
