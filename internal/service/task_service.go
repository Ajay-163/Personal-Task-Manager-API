package service

import (
	"errors"
	"personal-task-manager/internal/models"
	"personal-task-manager/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// function to create task service
func (s *TaskService) CreateTask(
	task models.Task,
) (models.Task, error) {

	if task.Title == "" {
		return models.Task{}, errors.New("title is required")
	}

	createdTask, err := s.repo.CreateTask(task)

	if err != nil {
		return models.Task{}, err
	}

	return createdTask, nil
}

// Get Task By ID Service
func (s *TaskService) GetTaskByID(id int) (models.Task, error) {

	task, found := s.repo.GetTaskByID(id)

	if !found {
		return models.Task{}, errors.New("task not found")
	}

	return task, nil
}

// Get lists of tasks service
func (s *TaskService) GetAllTasks() []models.Task {
	return s.repo.GetAllTasks()
}

// update task service
func (s *TaskService) UpdateTask(
	id int,
	task models.Task,
) (models.Task, error) {

	if task.Title == "" {
		return models.Task{}, errors.New("title is required")
	}

	updatedTask, found := s.repo.UpdateTask(id, task)

	if !found {
		return models.Task{}, errors.New("task not found")
	}

	return updatedTask, nil
}

// delete task function
func (s *TaskService) DeleteTask(id int) error {

	deleted := s.repo.DeleteTask(id)

	if !deleted {
		return errors.New("task not found")
	}

	return nil
}

// marking of completed tasks
func (s *TaskService) MarkCompleted(id int) (models.Task, error) {

	task, found := s.repo.GetTaskByID(id)

	if !found {
		return models.Task{}, errors.New("task not found")
	}

	task.Completed = true

	updatedTask, _ := s.repo.UpdateTask(id, task)

	return updatedTask, nil
}
