package main

import "fmt"

func main() {
	age := 45
	fmt.Println(age <= 20)

	if age < 20 {
		fmt.Println("age is less than 20")
	} else if age > 40 {
		fmt.Println("age is greater than 40")
	}
}
