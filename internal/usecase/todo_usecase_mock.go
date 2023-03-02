package usecase

import (
	"ddd-to-do-list/internal/aggregate"

	"github.com/stretchr/testify/mock"
)

type TodoUsecaseMock struct {
	mock.Mock
}

func (m *TodoUsecaseMock) GetTodo() (todo aggregate.Todos, err error) {
	args := m.Called()

	return args.Get(0).(aggregate.Todos), args.Error(1)
}

func (m *TodoUsecaseMock) GetTodoByID(id uint64) (res aggregate.Todos, err error) {
	args := m.Called(id)

	return args.Get(0).(aggregate.Todos), args.Error(1)
}

func (m *TodoUsecaseMock) CreateTodo(activitGroupID int, title, priority string) (uint64, error) {
	args := m.Called(activitGroupID, title, priority)

	return args.Get(0).(uint64), args.Error(1)
}

func (m *TodoUsecaseMock) UpdateTodo(id uint64, activitGroupID, IsActive int, title, priority string) error {
	args := m.Called(id, activitGroupID, IsActive, title, priority)

	return args.Error(0)
}

func (m *TodoUsecaseMock) DeleteTodo(id uint64) error {
	args := m.Called(id)

	return args.Error(0)
}
