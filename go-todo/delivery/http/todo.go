package delivery_http

import (
	"go-todo/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeliveryHTTP struct {
	usecase *usecase.Usecase
}

func NewHTTP(usecase *usecase.Usecase) *DeliveryHTTP {
	return &DeliveryHTTP{usecase: usecase}
}

func (d *DeliveryHTTP) ListTodos(c *gin.Context) {
	todos, err := d.usecase.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve todos"})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{"todos": todos})
}

func (d *DeliveryHTTP) CreateTodo(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty"})
		return
	}

	if err := d.usecase.AddTodo(title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create todo"})
		return
	}

	d.ListTodos(c)
}

func (d *DeliveryHTTP) ToggleTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := d.usecase.ToggleTodoStatus(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not toggle todo"})
		return
	}

	d.ListTodos(c)
}

func (d *DeliveryHTTP) DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := d.usecase.RemoveTodoByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete todo"})
		return
	}

	d.ListTodos(c)
}
