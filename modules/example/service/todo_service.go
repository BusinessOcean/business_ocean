package service

import (
	"bedatabase"
	"beservice/example/todo"
	"bytes"
	"context"
	"example/repository"
	"fmt"
)

type TodoService struct {
	repository repository.TodoRepository
	db         bedatabase.BeDatabase
	todo.UnimplementedTodoServiceServer
}

func NewTodoService(repository repository.TodoRepository, database bedatabase.BeDatabase) *TodoService {

	return &TodoService{repository: repository, db: database}
}

// CreateTodo implements todo.TodoServiceServer.
func (t *TodoService) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {

	t.db.Insert(req)

	return &todo.CreateTodoResponse{Todo: &todo.Todo{Title: req.Title, Id: 1827, Description: req.Description}}, nil
}

// DeleteTodo implements todo.TodoServiceServer.
func (t *TodoService) DeleteTodo(context.Context, *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {

	return &todo.DeleteTodoResponse{Success: true}, nil
}

// ReadTodo implements todo.TodoServiceServer.
func (t *TodoService) ReadTodo(context.Context, *todo.ReadTodoRequest) (*todo.ReadTodoResponse, error) {
	result, err := t.db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	return &todo.ReadTodoResponse{Todo: &todo.Todo{Title: createKeyValuePairs(*result)}}, nil
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
