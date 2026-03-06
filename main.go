package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)


func main() {

	var todos []string
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("1. Enter TODO\n2. See TODOs\n3. Mark TODO complete\n4. Quit\n")
		fmt.Print("Select option: ")
		scanner.Scan()

		if scanner.Text() == "1" {
			fmt.Print("Enter a TODO: ")
			scanner.Scan()
			todos = append(todos, scanner.Text())
		} else if scanner.Text() == "2" {
			fmt.Print("Here's the list:\n")
			for i, v := range todos {
				fmt.Printf("%d: %v\n", i+1, v)
			}
		} else if scanner.Text() == "3" {
			fmt.Print("What TODO is finished? ")
			scanner.Scan()
			number, _ := strconv.Atoi(scanner.Text())
			number -= 1
			todo := todos[number]
			todos = slices.Delete(todos, number, number+1)
			fmt.Printf("%v marked as complete\n", todo)
		} else if scanner.Text() == "4" {
			fmt.Print("Bye!")
			break
		} else {
			fmt.Print("Unrecognized. Enter correct command.\n")
		}
		fmt.Print("\n")
	}
}