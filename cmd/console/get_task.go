package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"
)

func getTask(reader *bufio.Reader) {

	printHeader("GET TASK")

	id := readInput(reader, "Enter Task ID : ")
	//validating input
	taskID, err := strconv.Atoi(id)

	if err != nil {

		fmt.Println("Invalid Task ID")

		pressEnter(reader)

		return
	}
	//calling API
	body, err := get(fmt.Sprintf("/tasks/%d", taskID))

	if err != nil {

		printError(err)

		pressEnter(reader)

		return
	}
	//json decoding
	var task Task

	err = json.Unmarshal(body, &task)

	if err != nil {

		printError(err)
		pressEnter(reader)

		return
	}
	printTask(task)

	pressEnter(reader)
}
