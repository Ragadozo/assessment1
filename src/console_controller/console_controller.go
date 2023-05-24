package console_controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Ragadozo/assessment1/src/roman_converter"
)

func readFromConsole() string {
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	fmt.Printf("\n")
	return choice
}

func firstChoice() {
	fmt.Println("You selected option 1: ")
	fmt.Println("Enter a Roman number:")
	input := readFromConsole()
	result, err := roman_converter.RomanToDecimal(input)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The result: %d\n", result)
	}

}

func secondChoice() {
	fmt.Println("You selected option 2: ")
	fmt.Println("Enter a Decimal number:")
	input := readFromConsole()
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input.")
	} else {
		result, err := roman_converter.DecimalToRoman(number)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("The result: %s\n", result)
		}
	}

}

func Run_program() {

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Write '1' to convert Roman number to Decimal")
		fmt.Println("2. Write '2' to convert Decimal number to Roman")
		fmt.Println("Enter your choice:")

		choice := readFromConsole()

		switch choice {

		case "1":
			firstChoice()
		case "2":
			secondChoice()
		default:
			fmt.Println("Invalid choice.")
		}

		fmt.Println("Press Enter to continue or type 'exit' to quit:")

		continueInput := readFromConsole()

		if continueInput == "exit" {
			break
		}
	}

	fmt.Println("Program ended.")
}
