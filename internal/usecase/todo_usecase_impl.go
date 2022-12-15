package usecase

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/repository"
	"errors"
)

type todoUsecase struct {
	repo repository.TodoRepository
}

func (u *todoUsecase) GetTodo() (aggregate.Todos, error) {
	todo, err := u.repo.GetTodo()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUsecase) GetTodoByID(id uint64) (aggregate.Todos, error) {
	todo, err := u.repo.GetTodoByID(id)
	if len(todo) <= 0 {
		return nil, errors.New("data not found")
	}
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUsecase) CreateTodo(activitGroupID int, title, priority string) (uint64, error) {
	id, err := u.repo.CreateTodo(activitGroupID, title, priority)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *todoUsecase) UpdateTodo(id uint64, activitGroupID, IsActive int, title, priority string) error {
	err := u.repo.UpdateTodo(id, activitGroupID, IsActive, title, priority)
	if err != nil {
		return err
	}

	return nil
}

func (u *todoUsecase) DeleteTodo(id uint64) error {
	err := u.repo.DeleteTodo(id)
	if err != nil {
		return err
	}

	return nil
}

func NewTodoUsecase(repo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		repo: repo,
	}
}
