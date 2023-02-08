package main

import "fmt"

func main() {
	// Arrays
	//var ages [3]int = [3]int{1, 2, 3}
	var ages = [3]int{1, 2, 3}
	names := [4]string{"yoshi", "mario", "peach", "bowser"} // shorthand
	names[1] = "shivansh"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 5, 55}
	scores[2] = 44
	fmt.Println(scores, len(scores))
	scores = append(scores, 77) // does not change the actual scores but only return a new slice
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:2]
	rangeTwo := scores[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Thakur")
	fmt.Println(rangeOne)

}
