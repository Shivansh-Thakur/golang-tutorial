package main

import "fmt"

// group A: int, string, array, floats
// pass by value therefore return is required
func updateValue(x string) string {
	x = "thakur"
	return x
}

// group B (Pointer Wrapper values): maps, slices, functions
// creates a new map pointing to the same old map values
// but having different map address
func updateMenu(m map[string]float64) {
	m["pie"] = 10
}

func main() {
	name := updateValue("shivansh")
	fmt.Println(name)

	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
