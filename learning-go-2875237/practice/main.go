package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number: ")
	inputNum, _ := reader.ReadString('\n')
	convertedFloat, err := strconv.ParseFloat(strings.TrimSpace(inputNum), 64)
	if err != nil {
		fmt.Println("Its an error:", err)
	} else {
		fmt.Println("The number entered is:", convertedFloat)
	}

}
