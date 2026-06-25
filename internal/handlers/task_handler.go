package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"personal-task-manager/internal/models"
	"personal-task-manager/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

// helper functions
func writeJSON(
	w http.ResponseWriter,
	status int,
	data interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// individual handlers
// POST tasks method
func (h *TaskHandler) CreateTask(
	w http.ResponseWriter,
	r *http.Request,
) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
		return
	}

	createdTask, err := h.service.CreateTask(task)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusCreated,
		createdTask,
	)
}

// GET tasks
func (h *TaskHandler) GetAllTasks(
	w http.ResponseWriter,
	r *http.Request,
) {

	tasks := h.service.GetAllTasks()

	writeJSON(
		w,
		http.StatusOK,
		tasks,
	)
}

// creating helper for extracting ID's from URL
func getIDFromPath(path string) (int, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) != 2 {
		return 0, errors.New("invalid path")
	}

	return strconv.Atoi(parts[1])
}

// GET/ tasks/{id}
func (h *TaskHandler) GetTaskByID(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := getIDFromPath(r.URL.Path)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid task id",
			},
		)
		return
	}

	task, err := h.service.GetTaskByID(id)

	if err != nil {
		writeJSON(
			w,
			http.StatusNotFound,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		task,
	)
}

// PUT/tasks/{id}
func (h *TaskHandler) UpdateTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := getIDFromPath(r.URL.Path)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid task id",
			},
		)
		return
	}

	var task models.Task

	err = json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
		return
	}

	updatedTask, err := h.service.UpdateTask(id, task)

	if err != nil {
		writeJSON(
			w,
			http.StatusNotFound,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		updatedTask,
	)
}

// DELETE/tasks/{id}
func (h *TaskHandler) DeleteTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := getIDFromPath(r.URL.Path)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid task id",
			},
		)
		return
	}

	err = h.service.DeleteTask(id)

	if err != nil {
		writeJSON(
			w,
			http.StatusNotFound,
			map[string]string{
				"error": err.Error(),
			},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "task deleted successfully",
		},
	)
}

// Route dispatchers
// Collection handler
func (h *TaskHandler) Tasks(
	w http.ResponseWriter,
	r *http.Request,
) {

	switch r.Method {

	case http.MethodPost:
		h.CreateTask(w, r)

	case http.MethodGet:
		h.GetAllTasks(w, r)

	default:
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

// Resource handler
func (h *TaskHandler) TaskByID(
	w http.ResponseWriter,
	r *http.Request,
) {

	switch r.Method {

	case http.MethodGet:
		h.GetTaskByID(w, r)

	case http.MethodPut:
		h.UpdateTask(w, r)

	case http.MethodDelete:
		h.DeleteTask(w, r)

	default:
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}
