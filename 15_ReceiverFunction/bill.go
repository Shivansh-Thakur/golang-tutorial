package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 45.5, "cake": 42.2},
		tip:   0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bill breakdown is:\n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v...$%v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-25v...$%0.2f", "total:", total)

	return fs
}
