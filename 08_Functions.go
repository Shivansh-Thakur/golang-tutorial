package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func saybye(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("shivansh")
	sayGreeting("sanyam")
	saybye("shivansh")

	cycleName([]string{"shaun", "mario", "bratt", "ujwal"}, sayGreeting)
	a := circleArea(2.5)
	fmt.Println(a)
}
