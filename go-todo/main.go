package main

import (
	"flag"
	"go-todo/config"
	delivery_grpc "go-todo/delivery/grpc"
	delivery_http "go-todo/delivery/http"
	"go-todo/proto"
	"go-todo/repository"
	"go-todo/usecase"
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatalf("DB_PATH environment variable not set")
	}

	// Initialize
	db := config.InitDB(dbPath)
	repo := repository.NewTodoRepository(db)
	uc := usecase.NewTodoUsecase(repo)
	deliveryHTTP := delivery_http.NewHTTP(uc)
	grpcServer := grpc.NewServer()
	deliveryGRPC := delivery_grpc.NewGRPC(uc)

	// Define the server type flag
	serverType := flag.String("server", "http", "Specify the server type: 'http' or 'grpc'")
	flag.Parse()

	switch *serverType {
	case "grpc":
		startGRPC(deliveryGRPC, grpcServer)
	case "http":
		startHTTP(deliveryHTTP)
	default:
		log.Fatalf("Unknown server type: %s", *serverType)
	}
}

func startGRPC(delivery *delivery_grpc.DeliveryGRPC, grpcServer *grpc.Server) {
	proto.RegisterTodoServiceServer(grpcServer, delivery)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func startHTTP(delivery *delivery_http.DeliveryHTTP) {
	// HTTP
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", delivery.ListTodos)
	router.POST("/todos", delivery.CreateTodo)
	router.POST("/todos/:id/toggle", delivery.ToggleTodo)
	router.POST("/todos/:id/delete", delivery.DeleteTodo)

	router.Run(":8080")
}
