package adapters

import (
	"bedatabase"
	"beservice/example/todo"
	"fmt"
)

type ExampleAdapter interface {
	GetTodoById(id string)
	GetTodos() (interface{}, error)
	CreateTodo(todo todo.Todo)
	DeleteTodo(id string)
}

func NewExampleAdapter(db bedatabase.BeDatabase) ExampleAdapter {
	return exampleAdapter{db}
}

type exampleAdapter struct {
	bedatabase.BeDatabase
}

func (e exampleAdapter) GetTodoById(id string) {

	e.Query(id)
}
func (e exampleAdapter) GetTodos() (interface{}, error) {

	e.Query("id")

	return nil, fmt.Errorf("no data found")
}
func (e exampleAdapter) CreateTodo(todo todo.Todo) {

	e.Insert(todo)

}
func (e exampleAdapter) DeleteTodo(id string) {
	e.Delete(823974)
}
