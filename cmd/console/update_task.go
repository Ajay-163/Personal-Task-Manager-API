package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func updateTask(reader *bufio.Reader) {

	printHeader("UPDATE TASK")

	id := readInput(reader, "Enter Task ID : ")

	taskID, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Invalid Task ID")
		pressEnter(reader)
		return
	}
	//get existing task
	body, err := get(fmt.Sprintf("/tasks/%d", taskID))

	if err != nil {
		printError(err)
		pressEnter(reader)
		return
	}

	var task Task

	json.Unmarshal(body, &task)

	fmt.Println()
	fmt.Println("Current Task")

	printTask(task)
	//reading new vakues
	task.Title = readInput(reader, "\nNew Title : ")

	task.Description = readInput(reader, "New Description : ")

	status := readInput(reader, "Completed (y/n): ")

	task.Completed = strings.EqualFold(status, "y")
	// put request
	body, err = put(
		fmt.Sprintf("/tasks/%d", taskID),
		task,
	)

	if err != nil {
		printError(err)
		pressEnter(reader)
		return
	}
	//decode response
	var updated Task

	json.Unmarshal(body, &updated)

	printHeader("TASK UPDATED")

	printTask(updated)

	pressEnter(reader)
}
