package repository

import (
	"beservice/example/todo"
	"example/adapters"
)

type TodoRepository interface {
	GetTodoById(id string)
	GetTodos() (interface{}, error)
	CreateTodo(todo todo.Todo)
	DeleteTodo(id string)
}

type todoRepository struct {
	adapters.ExampleAdapter
}

func NewTodoRepository(adapter adapters.ExampleAdapter) TodoRepository {
	return &todoRepository{adapter}
}
