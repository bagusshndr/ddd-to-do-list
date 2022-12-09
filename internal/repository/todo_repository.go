package repository

import "ddd-to-do-list/internal/aggregate"

type TodoRepository interface {
	GetTodo() (todo aggregate.Todos, err error)
	CreateTodo(todo *aggregate.Todo) error
	GetTodoByID(id uint64) (res aggregate.Todos, err error)
	DeleteTodo(id uint64) error
}
