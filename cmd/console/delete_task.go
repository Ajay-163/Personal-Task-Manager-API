package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"
)

func deleteTask(reader *bufio.Reader) {

	printHeader("DELETE TASK")

	id := readInput(reader, "Enter Task ID : ")

	taskID, err := strconv.Atoi(id)

	if err != nil {

		fmt.Println("Invalid Task ID")

		pressEnter(reader)

		return
	}
	body, err := get(fmt.Sprintf("/tasks/%d", taskID))

	if err != nil {

		printError(err)

		pressEnter(reader)

		return
	}

	var task Task

	json.Unmarshal(body, &task)

	printTask(task)
	answer := readInput(reader, "\nDelete this task? (y/n): ")

	if answer != "y" && answer != "Y" {

		fmt.Println("\nDeletion Cancelled")

		pressEnter(reader)

		return
	}
	err = deleteRequest(
		fmt.Sprintf("/tasks/%d", taskID),
	)

	if err != nil {

		printError(err)

		pressEnter(reader)

		return
	}
	fmt.Println("Reached after deleteRequest()")

	fmt.Println("Printing success header...")

	printHeader("TASK DELETED")

	fmt.Println("Task deleted successfully.")

	pressEnter(reader)
}
