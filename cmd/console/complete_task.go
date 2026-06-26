package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"
)

func completeTask(reader *bufio.Reader) {

	printHeader("MARK TASK COMPLETED")

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

	err = json.Unmarshal(body, &task)

	if err != nil {

		printError(err)

		pressEnter(reader)

		return
	}
	printTask(task)
	if task.Completed {

		fmt.Println()

		fmt.Println("Task is already completed.")

		pressEnter(reader)

		return
	}
	if task.Completed {

		fmt.Println()

		fmt.Println("Task is already completed.")

		pressEnter(reader)

		return
	}
	task.Completed = true
	body, err = put(
		fmt.Sprintf("/tasks/%d", taskID),
		task,
	)

	if err != nil {

		printError(err)

		pressEnter(reader)

		return
	}
	var updated Task

	err = json.Unmarshal(body, &updated)

	if err != nil {

		printError(err)
		pressEnter(reader)

		return
	}
	printHeader("TASK COMPLETED")

	printTask(updated)

	pressEnter(reader)
}
