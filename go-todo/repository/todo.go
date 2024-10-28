package repository

import (
	"database/sql"
	"fmt"
	"go-todo/models"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) FindAll() ([]models.Todo, error) {
	rows, err := r.db.Query("SELECT id, description, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *Repo) Create(title string) error {
	_, err := r.db.Exec("INSERT INTO todos (description, completed, created_at, updated_at) VALUES (?, false, ?, ?)", title, time.Now(), time.Now())
	if err != nil {
		fmt.Println("error when inserting", err)
		return err
	}
	return err
}

func (r *Repo) Toggle(id int) error {
	_, err := r.db.Exec("UPDATE todos SET completed = NOT completed, updated_at = ? WHERE id = ?", time.Now(), id)
	return err
}

func (r *Repo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
