package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
)

func createTask(reader *bufio.Reader) {

	var task Task

	/*fmt.Print("\nEnter Title: ")
	title, _ := reader.ReadString('\n')
	task.Title = strings.TrimSpace(title)*/

	task.Title = readInput(reader, "\nEnter Title : ")

	task.Description = readInput(reader, "Enter Description : ")

	fmt.Print("Enter Description: ")
	description, _ := reader.ReadString('\n')
	task.Description = strings.TrimSpace(description)

	task.Completed = false

	body, err := post("/tasks", task)

	if err != nil {
		printError(err)
		return
	}

	var createdTask Task

	err = json.Unmarshal(body, &createdTask)

	if err != nil {
		printError(err)
		return
	}

	printHeader("Task Created Successfully")

	fmt.Println("ID          :", createdTask.ID)
	fmt.Println("Title       :", createdTask.Title)
	fmt.Println("Description :", createdTask.Description)
	fmt.Println("Completed   :", createdTask.Completed)

	pressEnter(reader)
}
