package main

import "fmt"

func main() {
	age := 22
	name := "shivansh"
	// Print
	fmt.Print("hello")
	fmt.Print("world!") // no new line is added using Print
	fmt.Print("\nnew line\n")

	// Println
	fmt.Println("hello")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formated strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v\n", age, name)
	fmt.Printf("my age is %q and my name is %q\n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("age is of type %f \n", 22.55)
	fmt.Printf("age is of type %0.1f \n", 22.55) // rounded float value

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println(str)
}
