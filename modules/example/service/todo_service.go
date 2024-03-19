package service

import (
	"beservice/example/todo"
	"example/repository"
)

type TodoService struct {
	repository repository.TodoRepository
	todo.UnimplementedTodoServiceServer
}

func NewTodoService(repository repository.TodoRepository) *TodoService {
	return &TodoService{repository: repository}
}
