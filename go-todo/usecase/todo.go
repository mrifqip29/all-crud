package usecase

import (
	"go-todo/models"
	repositories "go-todo/repository"
)

type TodoService struct {
	repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	return s.repo.FindAll()
}

func (s *TodoService) AddTodo(title string) error {
	return s.repo.Create(title)
}

func (s *TodoService) ToggleTodoStatus(id int) error {
	return s.repo.Toggle(id)
}

func (s *TodoService) RemoveTodoByID(id int) error {
	return s.repo.Delete(id)
}
