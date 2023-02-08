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

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Please enter a valid price!")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item name", name, "Item price", price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Tip amount($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Please enter a valid tip!")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added is ", tip)
		promptOptions(b)
	case "s":
		fmt.Println("Your bill is saved", b)

	default:
		fmt.Println("That is not a valid option...")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
*/
