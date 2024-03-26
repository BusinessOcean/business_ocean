package service

import (
	"beservice/example/todo"
	"bytes"
	"context"
	"example/repository"
	"fmt"
)

type TodoService struct {
	repository repository.TodoRepository
	todo.UnimplementedTodoServiceServer
}

func NewTodoService(repository repository.TodoRepository) *TodoService {

	return &TodoService{repository: repository}
}

// CreateTodo implements todo.TodoServiceServer.
func (t *TodoService) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {

	t.repository.CreateTodo(todo.Todo{Title: req.Title, Description: req.Description})

	return &todo.CreateTodoResponse{Todo: &todo.Todo{Title: req.Title, Id: 1827, Description: req.Description}}, nil
}

// DeleteTodo implements todo.TodoServiceServer.
func (t *TodoService) DeleteTodo(context.Context, *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {

	return &todo.DeleteTodoResponse{Success: true}, nil
}

// ReadTodo implements todo.TodoServiceServer.
func (t *TodoService) ReadTodo(context.Context, *todo.ReadTodoRequest) (*todo.ReadTodoResponse, error) {
	result, err := t.repository.GetTodos()
	if err != nil {
		return nil, err
	}
	return &todo.ReadTodoResponse{Todo: &todo.Todo{Title: result.(string)}}, nil
}

// UpdateTodo implements todo.TodoServiceServer.
func (t *TodoService) UpdateTodo(context.Context, *todo.UpdateTodoRequest) (*todo.UpdateTodoResponse, error) {
	return &todo.UpdateTodoResponse{}, nil
}

func createKeyValuePairs(m map[interface{}]interface{}) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}
