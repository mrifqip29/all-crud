package usecase

import (
	"go-todo/models"
	repository "go-todo/repository"
)

type Usecase struct {
	repo *repository.Repo
}

func NewTodoUsecase(repo *repository.Repo) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) GetAllTodos() ([]models.Todo, error) {
	return u.repo.FindAll()
}

func (u *Usecase) AddTodo(title string) error {
	return u.repo.Create(title)
}

func (u *Usecase) ToggleTodoStatus(id int) error {
	return u.repo.Toggle(id)
}

func (u *Usecase) RemoveTodoByID(id int) error {
	return u.repo.Delete(id)
}
