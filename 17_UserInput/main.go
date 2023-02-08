package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getInput(promt string, r *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

/* uncomment
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)


	//	fmt.Print("Enter the bill Name:")
	//	name, _ := reader.ReadString('\n')
	//	name = strings.TrimSpace(name)


	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip)", reader)
	fmt.Println(opt)
}

func main() {
	mybill := createBill()
	promptOptions(mybill)

	fmt.Println(mybill)
}
*/
