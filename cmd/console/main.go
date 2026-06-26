package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println()
		fmt.Println("===================================")
		fmt.Println("     PERSONAL TASK MANAGER")
		fmt.Println("===================================")
		fmt.Println("1. Create Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Get Task By ID")
		fmt.Println("4. Update Task")
		fmt.Println("5. Delete Task")
		fmt.Println("6. Mark Task Completed")
		fmt.Println("7. Exit")
		fmt.Println()

		fmt.Print("Choose Option : ")

		choice, _ := reader.ReadString('\n')

		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":
			createTask(reader)
		case "2":
			listTasks(reader)
		case "3":
			getTask(reader)

		case "4":
			updateTask(reader)
		case "5":
			deleteTask(reader)
		case "6":
			completeTask(reader)

		case "7":
			fmt.Println("Thank you!")
			return

		default:
			fmt.Println("Invalid Option")
		}
	}
}
