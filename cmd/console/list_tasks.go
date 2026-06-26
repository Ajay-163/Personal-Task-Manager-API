package main

import (
	"bufio"
	"encoding/json"
	"fmt"
)

func listTasks(reader *bufio.Reader) {

	body, err := get("/tasks")

	if err != nil {
		printError(err)
		return
	}

	var tasks []Task

	err = json.Unmarshal(body, &tasks)

	if err != nil {
		printError(err)
		return
	}
	if len(tasks) == 0 {

		printHeader("TASK LIST")

		fmt.Println("No Tasks Found")

		pressEnter(reader)

		return
	}
	printHeader("TASK LIST")
	fmt.Printf(
		"%-5s %-25s %-12s\n",
		"ID",
		"TITLE",
		"COMPLETED",
	)

	fmt.Println("-----------------------------------------------")
	for _, task := range tasks {

		status := "No"

		if task.Completed {
			status = "Yes"
		}

		fmt.Printf(
			"%-5d %-25s %-12s\n",
			task.ID,
			task.Title,
			status,
		)
	}
	fmt.Println("-----------------------------------------------")

	pressEnter(reader)
}
