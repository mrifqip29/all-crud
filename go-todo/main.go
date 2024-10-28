package main

import (
	"go-todo/config"
	"log"
	"os"

	"go-todo/delivery"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatalf("DB_PATH environment variable not set")
	}

	db := config.InitDB(dbPath)
	todoHandler := delivery.NewTodoHandler(db)

	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", todoHandler.ListTodos)
	router.POST("/todos", todoHandler.CreateTodo)
	router.POST("/todos/:id/toggle", todoHandler.ToggleTodo)
	router.POST("/todos/:id/delete", todoHandler.DeleteTodo)

	router.Run(":8080")
}
