package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// 1. strings package
	greetings := "hello there friend"
	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.ReplaceAll(greetings, "hello", "hi"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Split(greetings, " "))
	fmt.Println(strings.Index(greetings, "e"))

	// The original string does not change,
	// as all the methods just return a new string and a copy is created
	fmt.Println("\nThe original string is : ", greetings)

	// 2. sort package

	ages := []int{3, 4, 5, 2, 7, 1, 9}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 7)
	fmt.Println("index is", index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println("\n", names)

	fmt.Println("index is", sort.SearchStrings(names, "peach"))
}
