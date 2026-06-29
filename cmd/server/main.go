package main

import (
	"log"
	"net/http"

	"personal-task-manager/internal/database"
	"personal-task-manager/internal/handlers"
	"personal-task-manager/internal/repository"
	"personal-task-manager/internal/service"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewMySQLTaskRepository(db)

	taskService := service.NewTaskService(repo)

	taskHandler := handlers.NewTaskHandler(taskService)

	http.HandleFunc("/tasks", taskHandler.Tasks)
	http.HandleFunc("/tasks/", taskHandler.TaskByID)

	// routes will go here

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
