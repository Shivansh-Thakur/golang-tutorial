package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} else {
		return initials[0], "_"
	}
}

func main() {
	fn1, sn1 := getInitials("shivansh thakur")
	fmt.Printf("first initial is %v and second initial is %v", fn1, sn1)
}
