package main

/* uncomment
func getInput(promt string, r *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

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

		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Tip amount($): ", reader)

		fmt.Println(tip)
	case "s":
		fmt.Println("You chose s")

	default:
		fmt.Println("That is not a valid option...")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)

	fmt.Println(mybill)
}
*/
