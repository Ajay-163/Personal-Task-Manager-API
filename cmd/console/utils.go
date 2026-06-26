package main

import (
	"bufio"
	"fmt"
	"strings"
)

func readInput(reader *bufio.Reader, prompt string) string {

	fmt.Print(prompt)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Input Error:", err)
		return ""
	}

	return strings.TrimSpace(input)
}
func printHeader(title string) {

	fmt.Println()

	fmt.Println("===================================")
	fmt.Println(title)
	fmt.Println("===================================")
}
func pressEnter(reader *bufio.Reader) {

	fmt.Print("\nPress ENTER to continue...")

	reader.ReadString('\n')
}
func printTask(task Task) {

	printHeader("TASK DETAILS")

	status := "No"

	if task.Completed {
		status = "Yes"
	}

	fmt.Println("ID          :", task.ID)
	fmt.Println("Title       :", task.Title)
	fmt.Println("Description :", task.Description)
	fmt.Println("Completed   :", status)
	fmt.Println("Created At  :", task.CreatedAt)
	fmt.Println("Updated At  :", task.UpdatedAt)
}
func printSuccess(message string) {

	fmt.Println()

	fmt.Println("✅", message)

	fmt.Println()
}
func printError(err error) {

	fmt.Println()

	fmt.Println("❌", err)

	fmt.Println()
}
