package main

import "fmt"

func main() {
	/*
		x := 0
		for x < 5 {
			fmt.Println("hello, world!")
			x++
		}
	*/
	names := []string{"yoshi", "mario", "peach", "bowser"}
	/*
		for i := 0; i < len(names); i++ {
			fmt.Printf("The value is %v \n", names[i])
		}
	*/
	for index, value := range names {
		fmt.Printf("The value at position %v is %v \n", index, value)
	}

	// if we do not want to use index but still want to declare a varial (index as _)
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
		value = "new string"
	}

	// the original string remains the same
	fmt.Println("\nthe original string is", names)
}
