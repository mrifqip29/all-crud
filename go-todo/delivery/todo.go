package delivery

import (
	"database/sql"
	"fmt"
	"go-todo/repository"
	"go-todo/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *usecase.TodoService
}

func NewTodoHandler(db *sql.DB) *TodoHandler {
	repo := repository.NewTodoRepository(db)
	service := usecase.NewTodoService(repo)
	return &TodoHandler{service: service}
}

func (h *TodoHandler) ListTodos(c *gin.Context) {
	todos, err := h.service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve todos"})
		return
	}

	fmt.Println(todos)

	c.HTML(http.StatusOK, "index.html", gin.H{"todos": todos})
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty"})
		return
	}

	if err := h.service.AddTodo(title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create todo"})
		return
	}

	h.ListTodos(c)
}

func (h *TodoHandler) ToggleTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.ToggleTodoStatus(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not toggle todo"})
		return
	}

	h.ListTodos(c)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.RemoveTodoByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete todo"})
		return
	}

	h.ListTodos(c)
}
