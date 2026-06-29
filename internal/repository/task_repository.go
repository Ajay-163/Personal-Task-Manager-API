package repository

import (
	"database/sql"
	"fmt"
	"personal-task-manager/internal/models"
)

type MySQLTaskRepository struct {
	db *sql.DB
}

func NewMySQLTaskRepository(db *sql.DB) *MySQLTaskRepository {
	return &MySQLTaskRepository{
		db: db,
	}
}

// function to creating task
func (r *MySQLTaskRepository) CreateTask(
	task models.Task,
) (models.Task, error) {

	query := `
	INSERT INTO tasks
	(title, description, completed)
	VALUES (?, ?, ?)
	`

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Description,
		task.Completed,
	)

	if err != nil {
		return models.Task{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return models.Task{}, err
	}

	task.ID = int(id)

	return task, nil
}

// function to get task by id
func (r *MySQLTaskRepository) GetTaskByID(
	id int,
) (models.Task, bool) {

	var task models.Task

	query := `
	SELECT
	id,
	title,
	description,
	completed,
	created_at,
	updated_at
	FROM tasks
	WHERE id = ?
	`
	err := r.db.QueryRow(
		query,
		id,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	fmt.Println("ID Received:", id)

	if err != nil {
		fmt.Println("Query Error:", err)
		return models.Task{}, false
	}

	fmt.Println("Task Found:", task)

	return task, true

}

// function to get list of tasks
func (r *MySQLTaskRepository) GetAllTasks() []models.Task {

	query := `
	SELECT
	id,
	title,
	description,
	completed,
	created_at,
	updated_at
	FROM tasks
	`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {

		var task models.Task

		rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		tasks = append(tasks, task)
	}

	return tasks
}

// function to update tasks
func (r *MySQLTaskRepository) UpdateTask(
	id int,
	task models.Task,
) (models.Task, bool) {

	query := `
	UPDATE tasks
	SET title=?,
	    description=?,
	    completed=?
	WHERE id=?
	`

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Description,
		task.Completed,
		id,
	)

	if err != nil {
		return models.Task{}, false
	}

	rows, _ := result.RowsAffected()

	if rows == 0 {
		return models.Task{}, false
	}

	task.ID = id

	return task, true
}

// function to delete tasks
func (r *MySQLTaskRepository) DeleteTask(id int) bool {

	query := `
	DELETE FROM tasks
	WHERE id = ?
	`

	result, err := r.db.Exec(
		query,
		id,
	)

	if err != nil {
		return false
	}

	rows, _ := result.RowsAffected()

	return rows > 0
}
