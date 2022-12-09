package repository

import "ddd-to-do-list/internal/aggregate"

type TodoRepository interface {
	GetTodo() (todo aggregate.Todos, err error)
	CreateTodo(activitGroupID uint64, title string) error
	GetTodoByID(id uint64) (res aggregate.Todos, err error)
	UpdateTodo(id, uint64, activity_group_id int, title string, is_active int, priority string) error
	DeleteTodo(id uint64) error
}
