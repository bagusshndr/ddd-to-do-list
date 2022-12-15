package usecase

import "ddd-to-do-list/internal/aggregate"

type TodoUsecase interface {
	GetTodo() (aggregate.Todos, error)
	GetTodoByID(id uint64) (aggregate.Todos, error)
	CreateTodo(activitGroupID int, title, priority string) (uint64, error)
	UpdateTodo(id uint64, activitGroupID, IsActive int, title, priority string) error
	DeleteTodo(id uint64) error
}
