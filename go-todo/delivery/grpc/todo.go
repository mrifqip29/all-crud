package delivery_grpc

import (
	"context"
	"go-todo/models"
	"go-todo/proto"
	"go-todo/usecase"
)

type DeliveryGRPC struct {
	proto.UnimplementedTodoServiceServer
	service *usecase.Usecase
}

func NewGRPC(service *usecase.Usecase) *DeliveryGRPC {
	return &DeliveryGRPC{service: service}
}

func (d *DeliveryGRPC) CreateTodo(ctx context.Context, req *proto.CreateTodoRequest) (*proto.CreateTodoResponse, error) {
	err := d.service.AddTodo(req.Title)
	if err != nil {
		return nil, err
	}
	todos, _ := d.service.GetAllTodos() // retrieve todos to get the latest inserted item
	return &proto.CreateTodoResponse{Todo: convertToProto(todos[len(todos)-1])}, nil
}

func (d *DeliveryGRPC) GetTodos(ctx context.Context, req *proto.GetTodosRequest) (*proto.GetTodosResponse, error) {
	todos, err := d.service.GetAllTodos()
	if err != nil {
		return nil, err
	}

	protoTodos := []*proto.Todo{}
	for _, todo := range todos {
		protoTodos = append(protoTodos, convertToProto(todo))
	}

	return &proto.GetTodosResponse{Todos: protoTodos}, nil
}

func (d *DeliveryGRPC) ToggleTodo(ctx context.Context, req *proto.ToggleTodoRequest) (*proto.ToggleTodoResponse, error) {
	err := d.service.ToggleTodoStatus(int(req.Id))
	if err != nil {
		return nil, err
	}

	todos, _ := d.service.GetAllTodos()
	var updatedTodo models.Todo
	for _, todo := range todos {
		if todo.ID == int(req.Id) {
			updatedTodo = todo
			break
		}
	}

	return &proto.ToggleTodoResponse{Todo: convertToProto(updatedTodo)}, nil
}

func (d *DeliveryGRPC) DeleteTodo(ctx context.Context, req *proto.DeleteTodoRequest) (*proto.DeleteTodoResponse, error) {
	err := d.service.RemoveTodoByID(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &proto.DeleteTodoResponse{Message: "Todo deleted successfully"}, nil
}

func convertToProto(todo models.Todo) *proto.Todo {
	return &proto.Todo{
		Id:          int32(todo.ID),
		Description: todo.Description,
		Completed:   todo.Completed,
	}
}
