package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"soup":   4.99,
		"pie":    7.99,
		"toffee": 6.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
	phonebook := map[int]string{
		15: "shivansh",
		14: "sanyam",
		64: "harshit",
	}

	fmt.Println(phonebook)
	phonebook[15] = "kuldeep"
	fmt.Println(phonebook[15])
}
