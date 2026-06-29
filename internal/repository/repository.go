package repository

import "personal-task-manager/internal/models"

type TaskRepository interface {
	CreateTask(task models.Task) (models.Task, error)

	GetTaskByID(id int) (models.Task, bool)

	GetAllTasks() []models.Task

	UpdateTask(id int, task models.Task) (models.Task, bool)

	DeleteTask(id int) bool
}
